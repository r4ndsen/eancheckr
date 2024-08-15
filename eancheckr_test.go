package eancheckr

import "testing"

func TestVerifyEan(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid ean 13",
			args: args{"4006381333931"},
			want: 4006381333931,
		},
		{
			name: "valid gtin 14",
			args: args{"4006381333931"},
			want: 4006381333931,
		},
		{
			name: "valid isbn 13",
			args: args{"978-3-650-40122-9"},
			want: 9783650401229,
		},
		{
			name: "with interruption",
			args: args{"'9.7,8 3 =650 /401229 '"},
			want: 9783650401229,
		},
		{
			name:    "invalid isbn",
			args:    args{"978-3-650-40122-99"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "with alpha characters",
			args:    args{"4006381333931a"},
			wantErr: true,
		},
		{
			name:    "with multi alpha characters",
			args:    args{"a4006381333931b"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyEan(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyEan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyEan() got = %v, want %v", got, tt.want)
			}
		})
	}
}
