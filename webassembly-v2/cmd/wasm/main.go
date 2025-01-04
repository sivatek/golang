package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid number of arguments passed",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		jsonOutputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOutputTextArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output text area",
			}
			return result
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("Unable to parse JSON. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		jsonOutputTextArea.Set("value", pretty)
		return nil
	})
	return jsonFunc
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan struct{})
}
