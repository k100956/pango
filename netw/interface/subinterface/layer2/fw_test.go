package layer2

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/version"
    "github.com/inwinstack/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        version version.Number
        vsys string
        importVsys string
        imports []string
        conf Entry
    }{
        {version.Number{7, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.2"}, Entry{
            Name: "ethernet1/1.2",
            Tag: 2,
            NetflowProfile: "netflow profile",
            Comment: "v1 basic",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwLayer2{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.conf.Comment, func(t *testing.T) {
            var err error
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err = ns.Set(tc.vsys, EthernetInterface, "ethernet1/1", VirtualWire, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(EthernetInterface, "ethernet1/1", VirtualWire, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                }
                if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
                if tc.importVsys != mc.Vsys {
                    t.Errorf("vsys: %q != %q", tc.importVsys, mc.Vsys)
                }
                if !reflect.DeepEqual(tc.imports, mc.Imports) {
                    t.Errorf("imports: %#v != %#v", tc.imports, mc.Imports)
                }
            }
        })
    }
}
