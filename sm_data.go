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
 var defaultSessionType = "IPV4"
 var defaultSscMode = "SSC_MODE_1"
 request := PMNConverter( stored_ambr_val, defaultSessionType,defaultSscMode)
 client.PMNSubscriberConfig(context.Background(), request)
}

func PMNConverter(ambrval Session_ambr, defaultSessionType string,defaultSscMode string) *protos.PMNSubscriberData {

        
        singleNssai := &models.Snssai{
                Sst:    1,
                Sd:     "000001",
        }

        sessionAmbr := &models.Ambr{
                Downlink: ambrval.Dl_ambr,
                Uplink:   ambrval.Ul_ambr,
        }

        pduSessionTypes := &models.PduSessionTypes{
                DefaultSessionType : defaultSessionType ,
                AllowedSessionTypes :        "IPV4V6" ,
        }

        arp := &models.Arp{
                PriorityLevel : 1,
                PreemptVuln  : "PREEMPTABLE" ,
                PreemptCap  :   "NOT_PREEMPT" ,

        }

        internal_5gQosProfile := &models.SubscribedDefaultQos{
                Internal_5Qi : 5 ,
                Arp  :        arp ,
                PriorityLevel : 1,
        }

        sscModes := &models.SscModes{
                DefaultSscMode : defaultSscMode ,
                AllowedSscModes :   "SSC_MODE_1", 
        }

        dnnConfigurations := &models.DnnConfiguration{
                PduSessionTypes : pduSessionTypes,
                Internal_5GQosProfile :        internal_5gQosProfile ,
                SessionAmbr : sessionAmbr ,
                SscModes : sscModes ,
        }

        smsd := &models.SessionManagementSubscriptionData {
                SingleNssai:  singleNssai,
                DnnConfigurations :            dnnConfigurations,
        }
        return &protos.PMNSubscriberData{
                PlmnSmData : smsd,
        }
}
