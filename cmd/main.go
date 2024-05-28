package main

import (
	"fmt"
	"log"

	"github.com/alex11prog/opentsdb-client/logger"
	"github.com/alex11prog/opentsdb-client/opentsdb"
)

const (
	host        = "127.0.0.1"
	port        = 4242
	dialTimeout = 3
)

func main() {
	opentsdbClient := opentsdb.NewClient(host, port, dialTimeout)
	defer opentsdbClient.Close()
	
	// Set metrics
	metric := &opentsdb.UniMetric{
		MetricName: "test.opentsdb",
		TimeStamp:  1713700000,
		Value:      3000,
		Tags: map[string]interface{}{
			"host": "localhost",
			"port": "4242",
		},
	}

	rr, err := opentsdbClient.Put([]*opentsdb.UniMetric{metric})
	if err != nil {
		logger.Logger.Debugf("errInfo: %+v", err)
	} else {
		logger.Logger.Debug("put.resp.info", rr.RespInfo)
	}

	// Get metrics
	query := &opentsdb.QueryRequestGet{
		Start:      "2015/08/09-00:00:00",
		End:        "2022/09/09-13:25:40",
		Aggregator: "none",
		MetricName: "test.opentsdb",
		GroupTagFilters: map[string]string{
			"entity": "entity",
		},
	}
	queryRsp, errRsp, err := opentsdbClient.QueryByGet(query)

	if err != nil {
		log.Fatal(err)
	} else if errRsp != nil {
		log.Fatal(*errRsp)
	}
	fmt.Println(*queryRsp)
}
