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
 request := PMNConverter( stored_ambr_val)
 client.PMNSubscriberConfig(context.Background(), request)
}

func PMNConverter(ambrval Session_ambr) *protos.PMNSubscriberData {

        
        singleNssai := &models.Snssai{
                Sst:    1,
                Sd:     "000001",
        }

        sessionAmbr := &models.Ambr{
                Downlink: ambrval.Dl_ambr,
                Uplink:   ambrval.Ul_ambr,
        }

        pduSessTypes := &models.InternalPduSessionType{
                PduSessTypes : "IPV4" ,
        }

        pduSessionTypes := &models.PduSessionTypes{
                DefaultSessionType : pduSessTypes ,
                AllowedSessionTypes :      nil ,
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
                SscModes  : "SSC_MODE_1" ,
        }

        sscModes := &models.SscModes{
                DefaultSscMode : sscMode ,
                AllowedSscModes :   nil, 
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
                PlmnSmData : map[string]smsd,
        }
}
