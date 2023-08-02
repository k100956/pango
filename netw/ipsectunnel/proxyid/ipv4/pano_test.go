package ipv4

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        conf Entry
    }{
        {"any proto", Entry{
            Name: "test any",
            Local: "local",
            Remote: "remote",
            ProtocolAny: true,
        }},
        {"number proto", Entry{
            Name: "test any",
            Local: "local",
            Remote: "remote",
            ProtocolNumber: 42,
        }},
        {"tcp proto", Entry{
            Name: "test any",
            Local: "local",
            Remote: "remote",
            ProtocolTcpLocal: 1,
            ProtocolTcpRemote: 2,
        }},
        {"udp proto", Entry{
            Name: "test any",
            Local: "local",
            Remote: "remote",
            ProtocolUdpLocal: 3,
            ProtocolUdpRemote: 4,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoIpv4{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.AddResp("")
            err := ns.Set("template", "", "tunnel", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("template", "", "tunnel", tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
