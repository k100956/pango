package auth

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
        {"no secret", Entry{
            Name: "one",
        }},
        {"with secret", Entry{
            Name: "two",
            Secret: "secret",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoAuth{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
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
