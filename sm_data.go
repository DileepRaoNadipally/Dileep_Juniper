package main

import (
    "context"
    "fmt"
    protos "magma/lte/cloud/go/protos"
    models "magma/lte/cloud/go/protos/models"
    "google.golang.org/grpc"
    //"github.com/go-openapi/swag"
    "log"
)

type Session_ambr struct {
        Dl_ambr string
        Ul_ambr string
}

type pdutype struct {
        pdutype string
}

type ssc_mode struct {
        sscmode string
}

func main() {
 fmt.Println("Hello client ...")

 opts := grpc.WithInsecure()
 cc, err := grpc.Dial("localhost:50051", opts)
 if err != nil {
  log.Fatal(err)
 }
 defer cc.Close()

 client := protos.NewPMNSubscriberConfigServicerClient(cc)
 stored_ambr_val := Session_ambr{"2000 Mbps", "1000 Mbps"}
 stored_pdu_type := pdutype{"IPV4V6"}
 stored_ssc_mode := ssc_mode{"SSC_MODE_1"}
 request := PMNConverter( stored_ambr_val,stored_pdu_type,stored_ssc_mode)
 client.PMNSubscriberConfig(context.Background(), request)
}

func PMNConverter(ambrval Session_ambr,pdutypeval pdutype sscModeVal ssc_mode) *protos.PMNSubscriberData {

        
        singleNssai := &models.Snssai{
                Sst:    1,
                Sd:     "000001",
        }

        sessionAmbr := &models.Ambr{
                Downlink: ambrval.Dl_ambr,
                Uplink:   ambrval.Ul_ambr,
        }

        pduSessTypes := &models.InternalPduSessionType{
                PduSessTypes : pdutypeval.pdutype ,
        }

        pduSessionTypes := &models.PduSessionTypes{
                DefaultSessionType : pduSessTypes ,
                AllowedSessionTypes :      pduSessTypes ,
        }

        preemptionCapability := &models.PreemptionCapability{}

        preemptionVulnerability  := &models.PreemptionVulnerability{}

        arp := &models.Arp{
                PriorityLevel : 1,
                PreemptVuln  : preemptionVulnerability ,
                PreemptCap  :   preemptionCapability ,

        }

        internal_5gQosProfile := &models.SubscribedDefaultQos{
                Internal_5Qi : 5 ,
                Arp  :        arp ,
                PriorityLevel : 1,
        }

        sscMode := &models.InternalSscMode{
                SscModes  : sscModeVal ,
        }

        sscModes := &models.SscModes{
                DefaultSscMode : sscMode ,
                AllowedSscModes :   sscMode, 
        }

        dnnConfigurations := &models.DnnConfiguration{
                PduSessionTypes : pduSessionTypes,
                Internal_5GQosProfile :        internal_5gQosProfile ,
                SessionAmbr : sessionAmbr ,
                SscModes : sscModes ,
        }

        smsd := &models.SessionManagementSubscriptionData {
                SingleNssai:  singleNssai,
                DnnConfigurations :            map[string]dnnConfigurations,
        }
        return &protos.PMNSubscriberData{
                PlmnSmData : smsd,
        }
}
