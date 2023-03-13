//go:build (js && wasm) || (js && wasm) || (js && wasm)

package main

import (
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"syscall/js"
)

func executeWithPromise(funcToExec func(args []js.Value) (string, error)) js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, outerArgs []js.Value) any {

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				result, err := funcToExec(outerArgs)

				if err != nil {
					jsonError(err, reject)
				} else {
					resolve.Invoke(result)
				}
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})

	return jsonFunc
}

func listOrganisations() js.Func {
	return executeWithPromise(func(args []js.Value) (string, error) {
		token := args[0].String()
		client := safetorun.New(token)
		organisations, err := client.ListOrganisations()

		if err != nil {
			return "", err
		} else {
			return ToJson(organisations.Items), nil
		}

	})
}

func deleteOrganisation() js.Func {
	return executeWithPromise(func(args []js.Value) (string, error) {
		token := args[0].String()
		organisationId := args[1].String()
		client := safetorun.New(token)
		_, err := client.DeleteOrganisationAndWait(organisationId)

		if err != nil {
			return "", err
		} else {
			return "Organisation removed", nil
		}
	})
}

func createOrganisation() js.Func {
	return executeWithPromise(func(args []js.Value) (string, error) {
		token := args[0].String()
		organisationId := args[1].String()
		organisationName := args[2].String()

		client := safetorun.New(token)
		organisation, err := client.CreateOrganisationAndWait(safetorun.CreateOrganisationRequest{
			OrganisationId:   organisationId,
			OrganisationName: organisationName},
		)

		if err != nil {
			return "", err
		} else {
			return ToJson(organisation), nil
		}
	})
}

func jsonError(err error, reject js.Value) {
	errorConstructor := js.Global().Get("Error")
	errorObject := errorConstructor.New(fmt.Sprintf("{}", err))
	reject.Invoke(errorObject)
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("listOrgs", listOrganisations())
	js.Global().Set("deleteOrg", deleteOrganisation())
	js.Global().Set("createOrg", createOrganisation())
	<-make(chan bool)
}
