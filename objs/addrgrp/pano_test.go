package addrgrp

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        dg string
        conf Entry
    }{
        {"test static no tags", "", Entry{
            Name: "one",
            Description: "my description",
            StaticAddresses: []string{"adr1", "adr2"},
        }},
        {"test static with tags", "", Entry{
            Name: "one",
            Description: "my description",
            StaticAddresses: []string{"adr1", "adr2"},
            Tags: []string{"tag1", "tag2"},
        }},
        {"test dynamic no tags", "dg1", Entry{
            Name: "one",
            Description: "my description",
            DynamicMatch: "'tag1' or 'tag2' and 'tag3'",
        }},
        {"test dynamic with tags", "dg2", Entry{
            Name: "one",
            Description: "my description",
            DynamicMatch: "'tag1' or 'tag2' and 'tag3'",
            Tags: []string{"tag1", "tag2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoAddrGrp{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.dg, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.dg, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}

