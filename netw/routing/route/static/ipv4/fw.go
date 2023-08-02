package ipv4

import (
    "fmt"
    "encoding/xml"

    "github.com/inwinstack/pango/util"
    "github.com/inwinstack/pango/version"
)


// FwIpv4 is the client.Network.StaticRoute namespace.
type FwIpv4 struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *FwIpv4) Initialize(con util.XapiClient) {
    c.con = con
}

// ShowList performs SHOW to retrieve a list of IPv4 routes.
func (c *FwIpv4) ShowList(vr string) ([]string, error) {
    c.con.LogQuery("(show) list of IPv4 routes")
    path := c.xpath(vr, nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// GetList performs GET to retrieve a list of IPv4 routes.
func (c *FwIpv4) GetList(vr string) ([]string, error) {
    c.con.LogQuery("(get) list of IPv4 routes")
    path := c.xpath(vr, nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given IPv4 route.
func (c *FwIpv4) Get(vr, name string) (Entry, error) {
    c.con.LogQuery("(get) IPv4 route %q", name)
    return c.details(c.con.Get, vr, name)
}

// Show performs SHOW to retrieve information for the given IPv4 route.
func (c *FwIpv4) Show(vr, name string) (Entry, error) {
    c.con.LogQuery("(show) IPv4 route %q", name)
    return c.details(c.con.Show, vr, name)
}

// Set performs SET to create / update one or more IPv4 routes.
func (c *FwIpv4) Set(vr string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if vr == "" {
        return fmt.Errorf("vr must be specified")
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given routes.
    d := util.BulkElement{XMLName: xml.Name{Local: "static-route"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) IPv4 routes: %v", names)

    // Set xpath.
    path := c.xpath(vr, names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the IPv4 routes.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit performs EDIT to create / update an IPv4 route.
func (c *FwIpv4) Edit(vr string, e Entry) error {
    var err error

    if vr == "" {
        return fmt.Errorf("vr must be specified")
    }

    _, fn := c.versioning()

    c.con.LogAction("(edit) IPv4 route %q", e.Name)

    // Set xpath.
    path := c.xpath(vr, []string{e.Name})

    // Edit the IPv4 route.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given IPv4 routes.
//
// IPv4 routes can be a string or an Entry object.
func (c *FwIpv4) Delete(vr string, e ...interface{}) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if vr == "" {
        return fmt.Errorf("vr must be specified")
    }

    names := make([]string, len(e))
    for i := range e {
        switch v := e[i].(type) {
        case string:
            names[i] = v
        case Entry:
            names[i] = v.Name
        default:
            return fmt.Errorf("Unknown type sent to delete: %s", v)
        }
    }
    c.con.LogAction("(delete) IPv4 routes: %v", names)

    // Remove IPv4 routes.
    path := c.xpath(vr, names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for this namespace struct **/

func (c *FwIpv4) versioning() (normalizer, func(Entry) (interface{})) {
    v := c.con.Versioning()

    if v.Gte(version.Number{8, 0, 0, ""}) {
        return &container_v3{}, specify_v3
    } else if v.Gte(version.Number{7, 1, 0, ""}) {
        return &container_v2{}, specify_v2
    } else {
        return &container_v1{}, specify_v1
    }
}

func (c *FwIpv4) details(fn util.Retriever, vr, name string) (Entry, error) {
    path := c.xpath(vr, []string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *FwIpv4) xpath(vr string, vals []string) []string {
    return []string{
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "network",
        "virtual-router",
        util.AsEntryXpath([]string{vr}),
        "routing-table",
        "ip",
        "static-route",
        util.AsEntryXpath(vals),
    }
}
