package service

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo    Repository
		hasher  Hasher
		token   Token
		clocker Clocker
	}

	tests := []struct {
		name string
		args args
		want *service
	}{
		{
			name: "ErrorPointerWhenRepoIsNil",
			args: args{},
			want: &service{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.repo, tt.args.hasher, tt.args.token, tt.args.clocker)
		})
	}
}
