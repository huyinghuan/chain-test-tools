package setting

import (
	"chain-api-imgo/config"
	"log"
	"testing"
)

func TestSetup(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				env: "dev",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Setup(tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(config.GetConfig())
		})
	}
}
