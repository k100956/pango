package syslog

import (
    "testing"
    "reflect"

    "github.com/inwinstack/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := getTests()

    mc := &testdata.MockClient{}
    ns := &PanoSyslog{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err := ns.Set("", "", "", "shared", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("", "", "", "shared", tc.conf.Name)
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
