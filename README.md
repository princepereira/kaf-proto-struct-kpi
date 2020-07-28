# Kaf-Protobuf-Struct-Analyser

$ cd Kaf-Protobuf-Struct-Analyser/protos
$ protoc --go_out=../pkg/pbprotos/ *.proto

// Start Kafka

// Run the Kafka Consumer

$ cd Kaf-Protobuf-Struct-Analyser/pkg/test-read
$ go run test_read.go


// Run the Kafka Producer

$ cd Kaf-Protobuf-Struct-Analyser/pkg/test-write
$ go run test_write.go


Results:
=======

Protos
==========


1. Number of messages sent            : 1000000
2. Start Time before marshalling first message  : 2020-07-27 11:17:33.804063065 +0000 UTC m=+0.031314066
3. End Time after submitting last message     : 2020-07-27 11:17:48.997587632 +0000 UTC m=+15.224939465
4. Messages reached and unmarshalled      : 2020-07-27 11:19:40.77604124 +0000 UTC m=+2141.910906237
5. Time taken between point 2 and point4    : 2 min 6.971978175 seconds = 126.97197


Struct
==========


1. Number of messages sent            : 1000000
2. Start Time before marshalling first message  : 2020-07-27 11:13:12.084987846 +0000 UTC m=+0.016674526 
3. End Time after submitting last message     : 2020-07-27 11:13:36.250558101 +0000 UTC m=+24.182244771
4. Messages reached and unmarshaled       : 2020-07-27 11:16:04.163901019 +0000 UTC m=+1925.298766023
5. Time taken between point 2 and point4    : 2min 52.078913173 seconds = 172.07891
