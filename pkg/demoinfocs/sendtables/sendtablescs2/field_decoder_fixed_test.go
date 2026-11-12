package sendtablescs2

import "testing"

// fixed8 is 8 raw bits: as a varint, 0xFF would also eat the next byte
// (CS2 update of 2026-11-11, m_recentKillQueue).
func TestFindDecoderFixed8ReadsRawBytes(t *testing.T) {
	f := field{varName: "m_recentKillQueue", varType: "uint8[8]", encoder: "fixed8", fieldType: newFieldType("uint8")}
	dec := findDecoder(&f)
	r := newReader([]byte{0xFF, 0x01, 0x80})

	for _, want := range []uint64{0xFF, 0x01, 0x80} {
		if got := dec(r); got != want {
			t.Fatalf("fixed8 = %v, want %v", got, want)
		}
	}
}

func TestFindDecoderUint8WithoutEncoderStaysVarint(t *testing.T) {
	f := field{varName: "m_x", varType: "uint8", fieldType: newFieldType("uint8")}

	if got := findDecoder(&f)(newReader([]byte{0xAC, 0x02})); got != uint64(300) {
		t.Fatalf("varint = %v, want 300", got)
	}
}
