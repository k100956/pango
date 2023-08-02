package imp

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
        {"v1 minimal with deny", version.Number{7, 0, 0, ""}, Entry{
            Name: "minimal",
            Enable: false,
            Action: ActionDeny,
        }},
        {"v1 community none", version.Number{7, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: true,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
                "7.7.0.0": true,
                "8.8.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginIgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeRemove,
            CommunityType: CommunityTypeNone,
            ExtendedCommunityType: CommunityTypeRemoveAll,
        }},
        {"v1 community remove all", version.Number{7, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeRemoveAll,
            ExtendedCommunityType: CommunityTypeNone,
        }},
        {"v1 community remove regex", version.Number{7, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeRemoveRegex,
            CommunityValue: "my allow community regex",
            ExtendedCommunityType: CommunityTypeAppend,
            ExtendedCommunityValue: "append value",
        }},
        {"v1 community append", version.Number{7, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeAppend,
            CommunityValue: AppendNoExport,
            ExtendedCommunityType: CommunityTypeOverwrite,
            ExtendedCommunityValue: "overwrite value",
        }},
        {"v1 community overwrite", version.Number{7, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeOverwrite,
            CommunityValue: "overwrite value",
            ExtendedCommunityType: CommunityTypeRemoveRegex,
            ExtendedCommunityValue: "remove regex",
        }},
        {"v2 minimal with deny", version.Number{8, 0, 0, ""}, Entry{
            Name: "minimal",
            Enable: false,
            Action: ActionDeny,
        }},
        {"v2 community none", version.Number{8, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: true,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchRouteTable: MatchRouteTableUnicast,
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
                "7.7.0.0": true,
                "8.8.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginIgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeRemove,
            CommunityType: CommunityTypeNone,
            ExtendedCommunityType: CommunityTypeRemoveAll,
        }},
        {"v2 community remove all", version.Number{8, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchRouteTable: MatchRouteTableMulticast,
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeRemoveAll,
            ExtendedCommunityType: CommunityTypeNone,
        }},
        {"v2 community remove regex", version.Number{8, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchRouteTable: MatchRouteTableBoth,
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeRemoveRegex,
            CommunityValue: "my allow community regex",
            ExtendedCommunityType: CommunityTypeAppend,
            ExtendedCommunityValue: "append value",
        }},
        {"v2 community append", version.Number{8, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeAppend,
            CommunityValue: AppendNoExport,
            ExtendedCommunityType: CommunityTypeOverwrite,
            ExtendedCommunityValue: "overwrite value",
        }},
        {"v2 community overwrite", version.Number{8, 0, 0, ""}, Entry{
            Name: "foo",
            Enable: false,
            UsedBy: []string{"one", "two"},
            MatchAsPathRegex: "as path regex",
            MatchCommunityRegex: "community regex",
            MatchExtendedCommunityRegex: "extended community regex",
            MatchMed: "match med",
            MatchAddressPrefix: map[string] bool{
                "5.5.0.0": true,
                "6.6.0.0": false,
            },
            MatchNextHop: []string{"nh1", "nh2"},
            MatchFromPeer: []string{"fp1", "fp2"},
            Action: ActionAllow,
            Dampening: "dampening",
            LocalPreference: "local pref",
            Med: "1234",
            Weight: 25,
            NextHop: "nexthop",
            Origin: OriginEgp,
            AsPathLimit: 26,
            AsPathType: AsPathTypeNone,
            CommunityType: CommunityTypeOverwrite,
            CommunityValue: "overwrite value",
            ExtendedCommunityType: CommunityTypeRemoveRegex,
            ExtendedCommunityValue: "remove regex",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoImp{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.Version = tc.version
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
