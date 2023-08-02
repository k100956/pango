package redist

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/testdata"
    "github.com/inwinstack/pango/version"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        version version.Number
        conf Entry
    }{
        {"v1 disable", version.Number{7, 0, 0, ""}, Entry{
            Name: "one",
            Enable: false,
            Metric: 42,
            SetOrigin: "origin",
            SetMed: "med",
            SetLocalPreference: "localpref",
            SetAsPathLimit: 1,
            SetCommunity: []string{"com1", "com2"},
            SetExtendedCommunity: []string{"ec1", "ec2"},
        }},
        {"v1 enable", version.Number{7, 0, 0, ""}, Entry{
            Name: "two",
            Enable: true,
            Metric: 42,
            SetOrigin: "origin",
            SetMed: "med",
            SetLocalPreference: "localpref",
            SetAsPathLimit: 2,
            SetCommunity: []string{"com1", "com2"},
            SetExtendedCommunity: []string{"ec1", "ec2"},
        }},
        {"v2 disable", version.Number{8, 0, 0, ""}, Entry{
            Name: "three",
            Enable: false,
            AddressFamily: AddressFamilyIpv4,
            Metric: 42,
            SetOrigin: "origin",
            SetMed: "med",
            SetLocalPreference: "localpref",
            SetAsPathLimit: 3,
            SetCommunity: []string{"com1", "com2"},
            SetExtendedCommunity: []string{"ec1", "ec2"},
        }},
        {"v2 enable", version.Number{8, 0, 0, ""}, Entry{
            Name: "four",
            Enable: true,
            AddressFamily: AddressFamilyIpv6,
            RouteTable: RouteTableBoth,
            Metric: 42,
            SetOrigin: "origin",
            SetMed: "med",
            SetLocalPreference: "localpref",
            SetAsPathLimit: 4,
            SetCommunity: []string{"com1", "com2"},
            SetExtendedCommunity: []string{"ec1", "ec2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoRedist{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err := ns.Set("tmpl", "", "vr", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("tmpl", "", "vr", tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                }
                if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
