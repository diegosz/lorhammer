package lora

import (
	"encoding/base64"
	"github.com/brocaar/lora-gateway-bridge/gateway"
	"lorhammer/src/model"
	"testing"
)

func TestHandlePacket(t *testing.T) {
	pullAckData := []byte{2, 235, 169, 4}

	err := HandlePacket(pullAckData)

	if err != nil {
		t.Fatal("Something,went,wrong,while,handling,packet")
	}

	pushAckData := []byte{2, 165, 210, 1}

	err = HandlePacket(pushAckData)

	if err != nil {
		t.Fatal("Something,went,wrong,while,handling,packet")
	}

	pullRespData := []byte{2, 0, 0, 3, 123, 34, 116, 120, 112, 107, 34, 58, 123, 34, 105, 109, 109, 101, 34, 58, 102, 97, 108, 115, 101, 44, 34, 116, 109, 115, 116, 34, 58, 49, 49, 50, 51, 52, 53, 54, 44, 34, 102, 114, 101, 113, 34, 58, 56, 54, 54, 46, 51, 52, 57, 56, 49, 50, 44, 34, 114, 102, 99, 104, 34, 58, 48, 44, 34, 112, 111, 119, 101, 34, 58, 49, 52, 44, 34, 109, 111, 100, 117, 34, 58, 34, 76, 79, 82, 65, 34, 44, 34, 100, 97, 116, 114, 34, 58, 34, 83, 70, 55, 66, 87, 49, 50, 53, 34, 44, 34, 99, 111, 100, 114, 34, 58, 34, 52, 47, 54, 34, 44, 34, 105, 112, 111, 108, 34, 58, 116, 114, 117, 101, 44, 34, 115, 105, 122, 101, 34, 58, 49, 50, 44, 34, 100, 97, 116, 97, 34, 58, 34, 89, 76, 72, 107, 89, 86, 48, 103, 65, 65, 67, 100, 103, 55, 118, 122, 34, 125, 125}

	err = HandlePacket(pullRespData)

	if err != nil {
		t.Fatal("Something,went,wrong,while,handling,packet")
	}

	pullAckData = []byte{0, 0, 0, 4}

	err = HandlePacket(pullAckData)

	if err == nil {
		t.Fatal("An error should have been thrown for non-valid data array")
	}

	pushAckData = []byte{0, 0, 0, 1}

	err = HandlePacket(pushAckData)

	if err == nil {
		t.Fatal("An error should have been thrown for non-valid data array")
	}

	pullRespData = []byte{2, 0, 0, 3, 123}

	err = HandlePacket(pullRespData)

	if err == nil {
		t.Fatal("An error should have been thrown for non-valid data array")
	}

}

func TestPacket_Prepare(t *testing.T) {
	rxpks := make([]gateway.RXPK, 1)
	data := []byte{2, 165, 210, 1}
	rxpk := NewRxpk(data)
	rxpks[0] = rxpk
	gw := &model.Gateway{
		NsAddress:  "127.0.0.1",
		MacAddress: RandomEUI(),
	}
	packet, err := Packet{Rxpk: rxpks}.Prepare(gw)

	if err != nil {
		t.Fatal("An error occured")
	}

	if packet[0] != gateway.ProtocolVersion2 {
		t.Fatal("first byte should represent ProtocolVersion2")
	}

	if packet[3] != byte(gateway.PushData) {
		t.Fatal("first byte should represent ProtocolVersion2")
	}

}

func TestNewRxpk(t *testing.T) {
	data := []byte{2, 165, 210, 1}
	rxpk := NewRxpk(data)

	if int(rxpk.Size) != len(data) {
		t.Fatalf("Size parameter should represent the length of the data sent, found  %d  expected %d", rxpk.Size, len(data))
	}

	strData := base64.StdEncoding.EncodeToString(data)

	if rxpk.Data != strData {
		t.Fatal("Data parameter should represent the base64 encoding of data given")
	}

}