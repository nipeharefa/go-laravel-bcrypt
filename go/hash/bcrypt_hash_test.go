package hash

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		plain  string
		hashed string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test invalid",
			args: args{
				plain:  "123456",
				hashed: "$2y$10$G5kswFi9VW0YQExToOLpi.mSPFZElOofQ6yOXrKGytUT5rm9Qgeh2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Check(tt.args.plain, tt.args.hashed); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
