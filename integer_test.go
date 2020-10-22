package safecast

import "testing"

func TestInt32(t *testing.T) {
	tests := []struct {
		value int
		wantInt32 int32
		wantErr bool
	}{
		{
			value: 2147483647,
			wantInt32: 2147483647,
			wantErr: false,
		},
		{
			value: -2147483648,
			wantInt32: -2147483648,
			wantErr: false,
		},
		{
			value: 2147483648,
			wantErr: true,
		},
		{
			value: -2147483649,
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := Int32(test.value)
		if (test.wantErr && err == nil) || !test.wantErr && err != nil {
			t.Errorf("err is wrong, value=%v, wantErr=%v, err=%v", test.value, test.wantErr, err)
			continue
		}
		if got != test.wantInt32 {
			t.Errorf("value is wrong, value=%v, wantInt32=%v, got=%v", test.value, test.wantInt32, got)
		}
	}
}
