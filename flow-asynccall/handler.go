package function

import (
	"fmt"
	faasflow "github.com/s8sg/faas-flow"
)

// Define provide definiton of the workflow
func Define(flow *faasflow.Workflow, context *faasflow.Context) (err error) {
          flow.
        Apply("func1").
	Apply("func2").
        Modify(func(data []byte) ([]byte, error) {
	        // Do something
               	return data
        }).
        Callback("storage.io/bucket?id=3345612358265349126&file=" + context.Query.Get("filename")).
        OnFailure(func(err error) {
              // failure handler
        }).
        Finally(func(state string) {
              // cleanup code
        })
	
	return nil
}

// DefineStateStore provides the override of the default StateStore
func DefineStateStore() (faasflow.StateStore, error) {
	return nil, nil
}

// ProvideDataStore provides the override of the default DataStore
func DefineDataStore() (faasflow.DataStore, error) {
	return nil, nil
}
