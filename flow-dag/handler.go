package function

import (
	"fmt"
	faasflow "github.com/s8sg/faas-flow"
)

// Define provide definiton of the workflow
func Define(flow *faasflow.Workflow, context *faasflow.Context) (err error) {
	     dag := faasflow.CreateDag()
     dag.AddModifier("mod1", func(data []byte) ([]byte, error) {
     		// do something
		return data, nil
     })
     dag.AddFunction("call-f1", "func1")
     dag.AddFunction("call-f2", "func2")
     dag.AddModifier("mod2", func(data []byte) ([]byte, error) {
     		// do something
		return data, nil
     })
     // To Serialize multiple input the dag need be defined with a Aggregator
     dag.AddVertex("callback", faasflow.Aggregator(func(inputs map[string][]byte) ([]byte, error) {
				          mod2Data := inputs["mod2"]
					  func2Data := inputs["func2"]
				          // Serialize input for callback
					  return data, nil
				    }))
     dag.AddCallback("callback", "storage.io/bucket?id=3345612358265349126&file=" + context.Query.Get("filename"))
				    

     dag.AddEdge("mod1", "func1")
     dag.AddEdge("mod1", "func2")
     dag.AddEdge("func1", "mod2")
     dag.AddEdge("func2", "callback")
     dag.AddEdge("mod2", "callback")
     
     flow.ExecuteDag(dag)
}

// DefineStateStore provides the override of the default StateStore
func DefineStateStore() (faasflow.StateStore, error) {
	return nil, nil
}

// ProvideDataStore provides the override of the default DataStore
func DefineDataStore() (faasflow.DataStore, error) {
	return nil, nil
}
