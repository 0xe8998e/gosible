/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 16:44:06
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/gosible.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */

package gosible

import (
	"io/ioutil"
	"os"

	"github.com/flosch/pongo2/v5"
	"github.com/pkg/errors"
)

func TemplateString(template string, varsMap map[string]interface{}) (string, error) {

	tpl, err := pongo2.FromString(template)
	if err != nil {
		return " ", errors.Wrap(err, "pongo2.FromString Error")
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	ctxt := pongo2.Context{}
	ctxt = ctxt.Update(varsMap)

	out, err := tpl.Execute(ctxt)
	if err != nil {

		return " ", errors.Wrap(err, "tpl.Execute Error")
	}

	return out, nil
}

func TemplateFile(template string, varsMap map[string]interface{}) (string, error) {

	tpl, err := pongo2.FromFile(template)

	if err != nil {
		return "", errors.Wrap(err, "pongo2.FromFile not found")
	}

	// context, _ := tpl.ExecuteBytes(pctx)
	ctxt := pongo2.Context{}
	ctxt = ctxt.Update(varsMap)

	out, err := tpl.Execute(ctxt)

	if err != nil {
		return "", errors.Wrap(err, "tpl.Execute Error")
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "gosible-")
	if err != nil {
		return "", errors.Wrap(err, "Cannot create temporary file")
	}

	ioutil.WriteFile(tmpFile.Name(), []byte(out), 0644)

	return tmpFile.Name(), nil
}
