package keylog_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/wader/fq/format/tls/keylog"
)

func TestParse(t *testing.T) {
	b, err := os.ReadFile("testdata/keylog.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := keylog.Parse(string(b))
	if err != nil {
		t.Fatal(err)
	}

	expected := keylog.Map{
		keylog.Entry{
			Label: keylog.ClientRandom,
			ClientRandom: [32]uint8{
				0x52, 0x36, 0x2c, 0x10, 0x12, 0xcf, 0x23, 0x62, 0x82, 0x56, 0xe7, 0x45, 0xe9, 0x03, 0xce, 0xa6,
				0x96, 0xe9, 0xf6, 0x2a, 0x60, 0xba, 0x0a, 0xe8, 0x31, 0x1d, 0x70, 0xde, 0xa5, 0xe4, 0x19, 0x49,
			},
		}: []uint8{
			0x9f, 0x9a, 0x0f, 0x19, 0xa0, 0x2b, 0xdd, 0xbe, 0x1a, 0x05, 0x92, 0x65, 0x97, 0xd6, 0x22, 0xcc,
			0xa0, 0x6d, 0x2a, 0xf4, 0x16, 0xa2, 0x8a, 0xd9, 0xc0, 0x31, 0x63, 0xb8, 0x7f, 0xf1, 0xb0, 0xc6,
			0x78, 0x24, 0xbb, 0xdb, 0x59, 0x5b, 0x32, 0xd8, 0x02, 0x7d, 0xb5, 0x66, 0xec, 0x04, 0xfb, 0x25,
		},
		keylog.Entry{
			Label: keylog.ClientRandom,
			ClientRandom: [32]uint8{
				0x52, 0x36, 0x2c, 0x10, 0xa2, 0x66, 0x5e, 0x32, 0x3a, 0x2a, 0xdb, 0x4b, 0x9d, 0xa0, 0xc1, 0xd,
				0x4a, 0x88, 0x23, 0x71, 0x92, 0x72, 0xf8, 0xb4, 0xc9, 0x7a, 0xf2, 0x4f, 0x92, 0x78, 0x48, 0x12,
			},
		}: []uint8{
			0x9f, 0x9a, 0x0f, 0x19, 0xa0, 0x2b, 0xdd, 0xbe, 0x1a, 0x05, 0x92, 0x65, 0x97, 0xd6, 0x22, 0xcc,
			0xa0, 0x6d, 0x2a, 0xf4, 0x16, 0xa2, 0x8a, 0xd9, 0xc0, 0x31, 0x63, 0xb8, 0x7f, 0xf1, 0xb0, 0xc6,
			0x78, 0x24, 0xbb, 0xdb, 0x59, 0x5b, 0x32, 0xd8, 0x02, 0x7d, 0xb5, 0x66, 0xec, 0x04, 0xfb, 0x25,
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %+#v, got %+#v", expected, actual)
	}
}

func TestParseErr(t *testing.T) {
	_, err := keylog.Parse("nonsense")
	if err != nil {
		t.Errorf("should ignore nonsense")
	}
}
func TestParseErr2(t *testing.T) {
	_, err := keylog.Parse("CLIENT_RANDOM 1234 1234")
	if err == nil {
		t.Errorf("expected CLIENT_RANDOM parse error")
	}
}
