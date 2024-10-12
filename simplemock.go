package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

type templData struct {
	PackageName   string
	Imports       []*packages.Package
	InterfaceName string
	MockName      string
	Methods       []*types.Func
}

var (
	templ = template.Must(template.New("MockSource").Funcs(template.FuncMap{
		"typeString": typeString,
		"signature":  signature,
		"args":       defaultedArgs,
	}).Parse(
		`// Code generated by simplemock. DO NOT EDIT.
package {{ .PackageName }}
{{- with .Imports }}

import (
{{- range . }}
	"{{ .PkgPath }}"
{{- end }}
)

{{- end }}

var _ {{ .InterfaceName }} = &{{ .MockName }}{}

{{- $mockName := .MockName }}
{{ with .Methods }}
type {{ $mockName }} struct {
{{- range . }}
	{{ .Name }}Func {{ typeString .Type }}
{{- end }}
}
{{- else }}
type {{ $mockName }} struct {}
{{- end }}
{{- range .Methods }}

func (m *{{ $mockName }}) {{ .Name }}{{ signature .Signature }} {
	if m.{{ .Name }}Func != nil {
		return m.{{ .Name }}Func({{ args .Signature }})
	}
	panic("{{ .Name }} called with nil {{ .Name }}Func!")
}
{{- end }}
`,
	))
)

func main() {
	if len(os.Args) != 3 {
		println("Usage: simplemock interface outfile")
		os.Exit(1)
	}

	typeName := os.Args[1]
	fname := os.Args[2]

	inputFile := os.Getenv("GOFILE")
	if inputFile == "" {
		println("Expected GOFILE environment variable to be set.")
		println("You should be using a //go:generate directive.")
		os.Exit(1)
	}

	cfg := packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		Tests: true,
	}

	pkgs, err := packages.Load(&cfg, "file="+inputFile)
	if err != nil {
		fmt.Printf("Failed to load input package and its dependencies: %v\n", err)
		os.Exit(1)
	}

	obj := pkgs[0].Types.Scope().Lookup(typeName)
	if obj == nil {
		pkgName, typeName, found := strings.Cut(typeName, ".")
		if !found {
			pkgName, typeName = typeName, pkgName
		}

		for _, pkg := range pkgs[0].Imports {
			if pkg.Name != pkgName {
				continue
			}
			obj = pkg.Types.Scope().Lookup(typeName)
			if obj != nil {
				pkgs = append(pkgs, pkg)
				break
			}
		}
		if obj == nil {
			fmt.Printf("Could not find type '%s'\n", typeName)
			os.Exit(1)
		}
	}

	if _, ok := obj.(*types.TypeName); !ok {
		fmt.Printf("Type %v is not a named type\n", obj)
		os.Exit(1)
	}

	if !types.IsInterface(obj.Type()) {
		fmt.Printf("Type %v is not an interface\n", obj)
		os.Exit(1)
	}

	if !obj.Exported() {
		fmt.Printf("Interface %v is not exported\n", obj)
		os.Exit(1)
	}

	imports, err := importsUsedBy(obj, pkgs)
	if err != nil {
		fmt.Printf("Failed getting %v imports: %v\n", obj, err)
		os.Exit(1)
	}

	iface := obj.Type().Underlying().(*types.Interface)
	methods := make([]*types.Func, iface.NumMethods())
	for i := range methods {
		methods[i] = iface.Method(i)
	}

	pkgName := pkgs[0].Name
	if pkgName != "main" {
		pkgName = strings.TrimSuffix(pkgs[0].Name, "_test") + "_test"
	}

	file := os.Stdout
	if fname != "os.Stdout" {
		file, err = os.Open(fname)
		if err != nil {
			fmt.Printf("Failed to create output file '%s': %v\n", fname, err)
			os.Exit(1)
		}
	}

	templ.Execute(file, templData{
		PackageName:   pkgName,
		Imports:       imports,
		InterfaceName: types.TypeString(obj.Type(), relativeTo),
		MockName:      obj.Name() + "Mock",
		Methods:       methods,
	})
}

func importsUsedBy(obj types.Object, pkgs []*packages.Package) ([]*packages.Package, error) {
	var pkg *packages.Package = nil
	for _, p := range pkgs {
		if p.Types == obj.Pkg() {
			pkg = p
		}
	}
	if pkg == nil {
		return nil, errors.New("owning package not found")
	}

	var id *ast.Ident = nil
	for ident, object := range pkg.TypesInfo.Defs {
		if object == obj {
			id = ident
			break
		}
	}
	if id == nil {
		return nil, errors.New("ast mapping not found")
	}

	typeSpec, ok := id.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return nil, errors.New("ast decl is not a type")
	}

	v := interfaceVisitor{info: pkg.TypesInfo, pkgs: make(map[*types.Package]struct{})}
	ast.Walk(&v, typeSpec)

	imports := make([]*packages.Package, 0, len(v.pkgs)+1)

	if pkg.Name != "main" {
		imports = append(imports, pkg)
	}

	for tPkg := range v.pkgs {
		for _, pPkg := range pkg.Imports {
			if tPkg == pPkg.Types {
				imports = append(imports, pPkg)
			}
		}
	}
	return imports, nil
}

type interfaceVisitor struct {
	info *types.Info
	pkgs map[*types.Package]struct{}
}

func (v *interfaceVisitor) Visit(node ast.Node) ast.Visitor {
	if id, ok := node.(*ast.Ident); ok {
		obj := v.info.Uses[id]
		if obj == nil {
			obj = v.info.Defs[id]
		}

		if obj != nil && obj.Pkg() != nil {
			v.pkgs[obj.Pkg()] = struct{}{}
		}
	}
	return v
}

func typeString(typ types.Type) string {
	return types.TypeString(typ, (*types.Package).Name)
}

func signature(sig *types.Signature) string {
	params := sig.Params()
	args := make([]string, params.Len())
	for i := range args {
		param := params.At(i)
		if name := params.At(i).Name(); name == "" {
			args[i] = fmt.Sprintf("arg%d", i)
		} else {
			args[i] = name
		}

		args[i] += " " + typeString(param.Type())
	}

	argsStr := strings.Join(args, ", ")
	if sig.Results().Len() != 0 {
		return fmt.Sprintf("(%s) %s", argsStr, typeString(sig.Results()))
	}
	return fmt.Sprintf("(%s)", argsStr)
}

func defaultedArgs(sig *types.Signature) string {
	params := sig.Params()
	args := make([]string, params.Len())
	for i := range args {
		if name := params.At(i).Name(); name == "" {
			args[i] = fmt.Sprintf("arg%d", i)
		} else {
			args[i] = name
		}
	}

	return strings.Join(args, ", ")
}

func relativeTo(pkg *types.Package) string {
	if pkg.Name() == "main" {
		return types.RelativeTo(pkg)(pkg)
	}
	return (*types.Package).Name(pkg)
}
