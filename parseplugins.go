package dotgo

import (
	"fmt"
	"io/ioutil"

	"go/build"
	"go/parser"
	"go/token"

	"github.com/johnaoss/dotgo/plugins"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

// Load returns the slice of plugins, along with an error if one occurred.
// Files that will be read in must be a part of the proper file structure for a plugin.
// That will be described at some point.
func Load(paths ...string) ([]plugins.Plugin, error) {
	if len(paths) == 0 {
		return plugins.GetPlugins(), nil
	}

	// Otherwise, we need to interpret the files
	values := make([]plugins.Plugin, 0)
	for _, filename := range paths {
		// Rev up the interpreter
		parse := interp.New(interp.Options{GoPath: build.Default.GOPATH})
		parse.Use(stdlib.Symbols)

		// We're going to want to read and evaluate the file.
		fbytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s, given: %w", filename, err)
		}
		fstring := string(fbytes)
		_, err = parse.Eval(fstring)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file %s, given: %w", filename, err)
		}

		// We now want to grab the function that returns the slice containing
		// the desired types. To do so, we first need the package name.
		pname, err := packageName(fstring)
		if err != nil {
			return nil, fmt.Errorf("failed to get package name of file %s, given: %w", filename, err)
		}

		// We want each plugin file to have a function called `GetPlugins`
		// that returns a slice of plugins.
		fn, err := parse.Eval(pname + ".GetPlugins")
		if err != nil {
			return nil, fmt.Errorf("file %s missing GetPlugins function, err: %w", filename, err)
		}

		pfunc := fn.Interface().(func() []plugins.Plugin)

		values = append(values, pfunc()...)
	}

	return values, nil
}

func packageName(contents string) (string, error) {
	fset := token.NewFileSet()

	// parse the go soure file, but only the package clause
	astFile, err := parser.ParseFile(fset, "", contents, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	if astFile.Name == nil {
		return "", fmt.Errorf("no package name found")
	}

	return astFile.Name.Name, nil
}
