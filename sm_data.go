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

type DefaultSessionType struct {
        PduSessTypes  string
}

type DefaultSscMode struct {
        sscModes  string
}

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
 var defaultSessionType = []DefaultSessionType{PduSessTypes:"IPV4"}
 var defaultSscMode = []DefaultSscMode{sscModes:"SSC_MODE_1"}
 request := PMNConverter( stored_ambr_val, defaultSessionType,defaultSscMode)
 client.PMNSubscriberConfig(context.Background(), request)
}

func PMNConverter(ambrval Session_ambr, defaultSessionType []DefaultSessionType,defaultSscMode []DefaultSscMode) *protos.PMNSubscriberData {

        defaultSessionType1 := []*models.InternalPduSessionType{}
          for _, value := range defaultSessionType {
                  temp := new(models.InternalPduSessionType)
                  temp.pduSessionTypes = value.pduSessionTypes
                  defaultSessionType1 = append(defaultSessionType, temp)
          }

          defaultSscMode1 := []*models.InternalSscMode{}
          for _, value := range defaultSscMode {
                  temp := new(models.InternalSscMode)
                  temp.sscModes = value.sscModes
                  defaultSscMode1 = append(defaultSscMode, temp)
          }
        
        singleNssai := &models.Snssai{
                sst:    1,
                sd:     "000001",
        }

        sessionAmbr := &models.Ambr{
                Downlink: ambrval.Dl_ambr,
                Uplink:   ambrval.Ul_ambr,
        }

        pduSessionTypes := &models.PduSessionTypes{
                DefaultSessionType : defaultSessionType1 ,
                AllowedSessionTypes :        "IPV4V6" 
        }

        arp := &models.Arp{
                PriorityLevel : 1,
                preemptVuln  : "PREEMPTABLE" ,
                preemptCap  :   "NOT_PREEMPT" ,

        }

        internal_5gQosProfile := &models.SubscribedDefaultQos{
                Internal_5qi : 5 ,
                Arp  :        arp ,
                PriorityLevel : 1,
        }

        sscModes := &models.SscModes{
                DefaultSscMode : defaultSscMode1 ,
                AllowedSscModes :   []string{"SSC_MODE_1","SSC_MODE_2","SSC_MODE_3"} 
        }

        dnnConfigurations := &models.DnnConfiguration{
                PduSessionTypes : pduSessionTypes,
                Internal_5gQosProfile :        internal_5gQosProfile ,
                SessionAmbr : sessionAmbr ,
                SscModes : sscModes 
        }

        smsd := &models.SessionManagementSubscriptionData {
                SingleNssai:  singleNssai
                DnnConfigurations :            dnnConfigurations
        }
        return &protos.PMNSubscriberData{
                plmnSmData: smsd,
        }
}