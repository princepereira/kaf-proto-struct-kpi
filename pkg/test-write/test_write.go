package main

import (
	"Kaf-Protobuf/pkg/message"
	"Kaf-Protobuf/pkg/pbproto"
	"confluent-kafka-go/kafka"
	"fmt"
	"os"
	"time"

	"encoding/json"

	"github.com/golang/protobuf/proto"
)

// 1 million
var totalMsgs = 1000000

// var totalMsgs = 1

func writeResult(result, fName string) {
	file, err := os.Create("/tmp/" + fName)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(result)
}

func produceProtos(p *kafka.Producer, count int) (time.Time, time.Time) {
	// Produce messages to topic (asynchronously)
	topic := "protobuf"

	// Build and initialize user message.
	msg := new(pbproto.Kpi)
	msg.KpiType = pbproto.KpiType_CPUInfo
	msg.ParentId = 1
	msg.ResourceId = 2
	msg.RaisedTs = 123456.78
	msg.ReportedTs = 123456.78
	msg.ResourceType = "CPU"
	additMsg := new(pbproto.AdditionalMsg)
	additMsg.InstanceId = 123
	additMsg.InstanceName = "CPUInfo"
	additMsg.Util = 78.9
	msg.AdditionalMessage = additMsg

	startTime := time.Now()

	var idx int32
	idx = 1
	for i := 1; i <= count; i++ {
		msg.ParentId = idx
		msg.ResourceId = idx
		idx++
		btes, _ := proto.Marshal(msg)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          btes,
		}, nil)
	}

	endTime := time.Now()

	return startTime, endTime
}

func produceStructs(p *kafka.Producer, count int) (time.Time, time.Time) {
	// Produce messages to topic (asynchronously)
	topic := "struct"

	// Build and initialize user message.
	msg := new(message.Kpi)
	msg.KpiType = message.KpiType_CPUInfo
	msg.RaisedTs = 123456.78
	msg.ReportedTs = 123456.78
	msg.ResourceType = "CPU"
	additMsg := new(message.AdditionalMsg)
	additMsg.InstanceId = 123
	additMsg.InstanceName = "CPUInfo"
	additMsg.Util = 78.9
	msg.AdditionalMessage = additMsg

	startTime := time.Now()

	var idx int32
	idx = 1
	for i := 1; i <= count; i++ {
		msg.ParentId = idx
		msg.ResourceId = idx
		idx++
		btes, _ := json.Marshal(msg)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          btes,
		}, nil)
	}

	endTime := time.Now()

	return startTime, endTime
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	protoStartTime, protoEndTime := produceProtos(p, totalMsgs)
	// structStartTime, structEndTime := produceStructs(p, totalMsgs)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	// Wait couple of seconds for delivery reports
	time.Sleep(2 * time.Second)

	var result string
	result = fmt.Sprintf("Dealing with protobufs ==> Number of messages sent :%d  , \nStart Time : %v , \nEnd Time : %v\n\n", totalMsgs, protoStartTime, protoEndTime)
	writeResult(result, "proto.txt")
	// result = fmt.Sprintf("Dealing with structs ==> Number of messages sent :%d  , \nStart Time : %v , \nEnd Time : %v\n\n", totalMsgs, structStartTime, structEndTime)
	// writeResult(result, "struct.txt")
}
