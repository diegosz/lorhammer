package lora

import (
	"lorhammer/src/model"
	"testing"

	"github.com/brocaar/lorawan"
)

func TestGetDevAddrFromDevEUI(t *testing.T) {
	eui := lorawan.EUI64{
		byte(1), byte(2), byte(3), byte(4),
		byte(5), byte(6), byte(7), byte(8),
	}

	devAddr := GetDevAddrFromDevEUI(eui)

	if devAddr.String() != "05060708" {
		t.Fatalf("DevAddr string value must equal %s ", "05060708")
	}

}

func TestNode_GetPushDataPayload(t *testing.T) {
	// , "01B501002919000006018403131313121244"
	node := NewNode("19842bd94743246b367c2e90942a1f73",
		"19842bd94743246b367c2e90942a1f774",
		"",
		[]model.Payload{
			model.Payload{Value: "01B501002919000006018403131313121233"},
			model.Payload{Value: "01B501002919000006018403131313121244"},
		},
		true,
	)

	fcnt := uint32(1)
	dataPayload, err := GetPushDataPayload(node, fcnt)
	if err != nil {
		t.Fatal("Couldn't get PushData payload")
	}
	phyPayload := lorawan.PHYPayload{}
	err = phyPayload.UnmarshalBinary(dataPayload)
	if err != nil {
		t.Fatal("Couldn't unmarshall PHYPayload Binary")
	}

	if phyPayload.MHDR.MType != lorawan.MType(lorawan.ConfirmedDataUp) {
		t.Fatal("Push data messages should always be of type ConfirmedDataType")
	}

	if phyPayload.MHDR.Major != lorawan.LoRaWANR1 {
		t.Fatal("Push data messages should always be of protocol LoraWan")
	}

	if len(phyPayload.MIC) != 4 {
		t.Fatal("An exactly 4 bytes MIC should be sent with the message")
	}

	_, ok := phyPayload.MACPayload.(*lorawan.MACPayload)
	if !ok {
		t.Fatal("the MacPayload should be of Type MACPayload ")
	}

}

func TestNode_GetPushDataPayloadWithoutRandomAccessOnPayloadArray(t *testing.T) {
	// , "01B501002919000006018403131313121244"
	node := NewNode("19842bd94743246b367c2e90942a1f73",
		"19842bd94743246b367c2e90942a1f774",
		"",
		[]model.Payload{
			model.Payload{Value: "01B501002919000006018403131313121233", Date: 1488931200},
			model.Payload{Value: "01B501002919000006018403131313121244", Date: 1488931201},
			model.Payload{Value: "01B501002919000006018403131313121233", Date: 1488931202},
		},
		false,
	)
	for index := 0; index < 3; index++ {
		fcnt := uint32(1)
		dataPayload, err := GetPushDataPayload(node, fcnt)
		if err != nil {
			t.Fatal("Couldn't get PushData payload")
		}
		phyPayload := lorawan.PHYPayload{}
		err = phyPayload.UnmarshalBinary(dataPayload)
		if err != nil {
			t.Fatal("Couldn't unmarshall PHYPayload Binary")
		}

		if phyPayload.MHDR.MType != lorawan.MType(lorawan.ConfirmedDataUp) {
			t.Fatal("Push data messages should always be of type ConfirmedDataType")
		}

		if phyPayload.MHDR.Major != lorawan.LoRaWANR1 {
			t.Fatal("Push data messages should always be of protocol LoraWan")
		}

		if len(phyPayload.MIC) != 4 {
			t.Fatal("An exactly 4 bytes MIC should be sent with the message")
		}

		_, ok := phyPayload.MACPayload.(*lorawan.MACPayload)
		if !ok {
			t.Fatal("the MacPayload should be of Type MACPayload ")
		}
	}

}

func TestNode_GetPushDataPayloadWithoutRandomAccessOnPayloadArrayAndReload(t *testing.T) {
	// , "01B501002919000006018403131313121244"
	node := NewNode("19842bd94743246b367c2e90942a1f73",
		"19842bd94743246b367c2e90942a1f774",
		"",
		[]model.Payload{
			model.Payload{Value: "01B501002919000006018403131313121233", Date: 1488931200},
			model.Payload{Value: "01B501002919000006018403131313121244", Date: 1488931201},
			model.Payload{Value: "01B501002919000006018403131313121233", Date: 1488931202},
		},
		false,
	)
	for index := 0; index < 5; index++ {
		fcnt := uint32(1)
		dataPayload, err := GetPushDataPayload(node, fcnt)
		if err != nil {
			t.Fatal("Couldn't get PushData payload")
		}
		phyPayload := lorawan.PHYPayload{}
		err = phyPayload.UnmarshalBinary(dataPayload)
		if err != nil {
			t.Fatal("Couldn't unmarshall PHYPayload Binary")
		}

		if phyPayload.MHDR.MType != lorawan.MType(lorawan.ConfirmedDataUp) {
			t.Fatal("Push data messages should always be of type ConfirmedDataType")
		}

		if phyPayload.MHDR.Major != lorawan.LoRaWANR1 {
			t.Fatal("Push data messages should always be of protocol LoraWan")
		}

		if len(phyPayload.MIC) != 4 {
			t.Fatal("An exactly 4 bytes MIC should be sent with the message")
		}

		_, ok := phyPayload.MACPayload.(*lorawan.MACPayload)
		if !ok {
			t.Fatal("the MacPayload should be of Type MACPayload ")
		}
	}

}

func TestNode_GetPushDataPayloadWithEmptyPayloadArray(t *testing.T) {
	// , "01B501002919000006018403131313121244"
	node := NewNode("19842bd94743246b367c2e90942a1f73",
		"19842bd94743246b367c2e90942a1f774",
		"",
		[]model.Payload{},
		true,
	)

	fcnt := uint32(1)
	dataPayload, err := GetPushDataPayload(node, fcnt)
	if err != nil {
		t.Fatal("Couldn't get PushData payload")
	}
	phyPayload := lorawan.PHYPayload{}
	err = phyPayload.UnmarshalBinary(dataPayload)
	if err != nil {
		t.Fatal("Couldn't unmarshall PHYPayload Binary")
	}

	if phyPayload.MHDR.MType != lorawan.MType(lorawan.ConfirmedDataUp) {
		t.Fatal("Push data messages should always be of type ConfirmedDataType")
	}

	if phyPayload.MHDR.Major != lorawan.LoRaWANR1 {
		t.Fatal("Push data messages should always be of protocol LoraWan")
	}

	if len(phyPayload.MIC) != 4 {
		t.Fatal("An exactly 4 bytes MIC should be sent with the message")
	}

	_, ok := phyPayload.MACPayload.(*lorawan.MACPayload)
	if !ok {
		t.Fatal("the MacPayload should be of Type MACPayload ")
	}

}

func TestNewJoinRequestPHYPayload(t *testing.T) {
	node := NewNode("",
		"",
		"",
		[]model.Payload{
			model.Payload{Value: "01B501002919000006018403131313121233"},
			model.Payload{Value: "01B501002919000006018403131313121244"},
		}, true)

	dataPayload := GetJoinRequestDataPayload(node)
	jrPayload := lorawan.PHYPayload{}
	jrPayload.UnmarshalBinary(dataPayload)

	if jrPayload.MHDR.MType != lorawan.MType(lorawan.JoinRequest) {
		t.Fatal("Join request messages should be of type JoinRequest")
	}

	if jrPayload.MHDR.Major != lorawan.LoRaWANR1 {
		t.Fatal("Join request messages should always be of protocol LoraWan")
	}

	_, ok := jrPayload.MACPayload.(*lorawan.JoinRequestPayload)
	if !ok {
		t.Fatal("the MacPayload should be of Type JoinRequestPayload")
	}

}
