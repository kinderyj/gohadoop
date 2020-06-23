package main

import (
    "log"
    "os"

    "github.com/hortonworks/gohadoop/hadoop_yarn"
    yarn_conf "github.com/hortonworks/gohadoop/hadoop_yarn/conf"
    "github.com/hortonworks/gohadoop/hadoop_yarn/yarn_client"
)

const (
    GOPATH               = "GOPATH"
    APPLICATIONMASTER_GO = "APPLICATIONMASTER_GO"
)

func main() {
    goPath := os.Getenv(GOPATH)
    applicationMasterLocation := os.Getenv(APPLICATIONMASTER_GO)
    if len(applicationMasterLocation) == 0 {
        applicationMasterLocation = goPath + "/src/github.com/hortonworks/gohadoop/hadoop_yarn/examples/dist_shell/applicationmaster.go"
    }

    // Create YarnConfiguration
    conf, _ := yarn_conf.NewYarnConfiguration()

    // Create YarnClient
    yarnClient, _ := yarn_client.CreateYarnClient(conf)

    // Create new application to get ApplicationSubmissionContext
    //_, asc, _ := yarnClient.CreateNewApplication()

    // Get nodes states
    yarnClient.GetNodesState()


    // Get GetApplicationReport
    applicationIdProto := hadoop_yarn.ApplicationIdProto{}
    id := int32(97)
    applicationIdProto.Id = &id
    clusterTimestamp := int64(1587465966658)
    applicationIdProto.ClusterTimestamp = &clusterTimestamp
    aplication, err := yarnClient.GetApplicationReport(&applicationIdProto)
    if err != nil {
        log.Println("the err is ", err)
    } else {
        log.Println("the aplication is", aplication)
        log.Println("====the id is ", *aplication.ApplicationId.Id)
        log.Println("====the User is ", *aplication.User)
        log.Println("====the ContainerResourceRequest is ", aplication.ResourceRequests)
    }

}
