package router

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        doDefaults bool
        imports []string
        conf Entry
    }{
        {"with defaults and import", "vsys2", true, []string{"one"}, Entry{
            Name: "one",
            Interfaces: []string{"ethernet1/1", "ethernet1/2"},
        }},
        {"no defaults or import", "", false, []string{"two"}, Entry{
            Name: "two",
            Interfaces: []string{"ethernet1/3", "ethernet1/4"},
        }},
        {"with raw fields", "vsys3", true, []string{"three"}, Entry{
            Name: "three",
            raw: map[string] string{
                "ecmp": "<ecmp>raw ecmp</ecmp>",
                "multicast": "<multicast><raw>multicast</raw></multicast>",
                "protocol": "<protocol><some><proto>field</proto></some></protocol>",
                "routing": "<routing-table><route1>something</route1><route2>b</route2></routing-table>",
            },
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwRouter{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            if tc.doDefaults {
                tc.conf.Defaults()
            }
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.vsys, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                }
                if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
                if tc.vsys != mc.Vsys {
                    t.Errorf("vsys: %s != %s", tc.vsys, mc.Vsys)
                }
                if !reflect.DeepEqual(tc.imports, mc.Imports) {
                    t.Errorf("imports: %#v != %#v", tc.imports, mc.Imports)
                }
            }
        })
    }
}
