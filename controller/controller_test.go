package controller

import (
	"testing"

	"github.com/armosec/armotoy/model"
)

func TestYamlAlerts(t *testing.T) {
	ctrler, _ := InitController("../myres.json", "JSON", "v2")
	_, table, _ := ctrler.CreateResourcePage([]string{}, []string{"C-0057"})

	ref := table.GetCell(1, 0).GetReference()
	item, _ := ref.(*model.ResourceReference)

	ctrler.YAMLInspect(item, []string{"C-0057"})

}
