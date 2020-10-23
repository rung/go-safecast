package safecast

import "testing"

func TestAtoi32(t *testing.T) {
	tests := []struct {
		value string
		wantInt32 int32
		wantErr bool
	}{
		{
			value: "2147483647",
			wantInt32: 2147483647,
			wantErr: false,
		},
		{
			value: "-2147483648",
			wantInt32: -2147483648,
			wantErr: false,
		},
		{
			value: "2147483648",
			wantErr: true,
		},
		{
			value: "-2147483649",
			wantErr: true,
		},
		{
			value: "test",
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := Atoi32(test.value)
		if (test.wantErr && err == nil) || !test.wantErr && err != nil {
			t.Errorf("err is wrong, value=%v, wantErr=%v, err=%v", test.value, test.wantErr, err)
			continue
		}
		if got != test.wantInt32 {
			t.Errorf("value is wrong, value=%v, wantInt32=%v, got=%v", test.value, test.wantInt32, got)
		}
	}
}

func TestAtoi16(t *testing.T) {
	tests := []struct {
		value string
		wantInt16 int16
		wantErr bool
	}{
		{
			value: "32768",
			wantInt16: 32767,
			wantErr: false,
		},
		{
			value: "-32768",
			wantInt16: -32768,
			wantErr: false,
		},
		{
			value: "32768",
			wantErr: true,
		},
		{
			value: "-32769",
			wantErr: true,
		},
		{
			value: "test",
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := Atoi16(test.value)
		if (test.wantErr && err == nil) || !test.wantErr && err != nil {
			t.Errorf("err is wrong, value=%v, wantErr=%v, err=%v", test.value, test.wantErr, err)
			continue
		}
		if got != test.wantInt16 {
			t.Errorf("value is wrong, value=%v, wantInt16=%v, got=%v", test.value, test.wantInt16, got)
		}
	}
}

func TestAtoi8(t *testing.T) {
	tests := []struct {
		value string
		wantInt8 int8
		wantErr bool
	}{
		{
			value: "127",
			wantInt8: 127,
			wantErr: false,
		},
		{
			value: "-128",
			wantInt8: -128,
			wantErr: false,
		},
		{
			value: "128",
			wantErr: true,
		},
		{
			value: "-129",
			wantErr: true,
		},
		{
			value: "test",
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := Atoi8(test.value)
		if (test.wantErr && err == nil) || !test.wantErr && err != nil {
			t.Errorf("err is wrong, value=%v, wantErr=%v, err=%v", test.value, test.wantErr, err)
			continue
		}
		if got != test.wantInt8 {
			t.Errorf("value is wrong, value=%v, wantInt8=%v, got=%v", test.value, test.wantInt8, got)
		}
	}
}