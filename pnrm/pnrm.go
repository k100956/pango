/*
Package pnrm is the client.Panorama namespace.
*/
package pnrm


import (
    "github.com/inwinstack/pango/util"

    "github.com/inwinstack/pango/pnrm/dg"
    "github.com/inwinstack/pango/pnrm/plugins/gcp/account"
    "github.com/inwinstack/pango/pnrm/plugins/gcp/gke/cluster"
    "github.com/inwinstack/pango/pnrm/plugins/gcp/gke/cluster/group"
    "github.com/inwinstack/pango/pnrm/template"
    "github.com/inwinstack/pango/pnrm/template/stack"
    "github.com/inwinstack/pango/pnrm/template/variable"
)


// Pnrm is the panorama.DeviceGroup namespace.
type Pnrm struct {
    DeviceGroup *dg.Dg
    GcpAccount *account.Account
    GkeCluster *cluster.Cluster
    GkeClusterGroup *group.Group
    Template *template.Template
    TemplateStack *stack.Stack
    TemplateVariable *variable.Variable
}

// Initialize is invoked on panorama.Initialize().
func (c *Pnrm) Initialize(i util.XapiClient) {
    c.DeviceGroup = &dg.Dg{}
    c.DeviceGroup.Initialize(i)

    c.GcpAccount = &account.Account{}
    c.GcpAccount.Initialize(i)

    c.GkeCluster = &cluster.Cluster{}
    c.GkeCluster.Initialize(i)

    c.GkeClusterGroup = &group.Group{}
    c.GkeClusterGroup.Initialize(i)

    c.Template = &template.Template{}
    c.Template.Initialize(i)

    c.TemplateStack = &stack.Stack{}
    c.TemplateStack.Initialize(i)

    c.TemplateVariable = &variable.Variable{}
    c.TemplateVariable.Initialize(i)
}
