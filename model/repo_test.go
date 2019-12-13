package model

import (
	"os"
	"path/filepath"
	"testing"
)

func init() {
	_ = os.Chdir(filepath.Join(os.Getenv("GOPATH"), "src", "github.com/growerlab/codev-svc"))
}

func TestInitRepo(t *testing.T) {
	type args struct {
		repoPath string
		name     string
	}
	tests := []struct {
		name    string
		args    args
		want    *Repo
		wantErr bool
	}{
		{
			name: "create repo",
			args: args{
				repoPath: ReposDir,
				name:     "moli",
			},
			want: &Repo{
				Path: ReposDir,
				Name: "moli",
				Branches: []*Branch{
					{
						Name: DefaultBranch,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitRepo("", tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("InitRepo() got = nil")
				return
			}
			t.Errorf("==== %+v", got.Branches[0].Name)
			return
			if got.Branches[0].Name != tt.want.Branches[0].Name {
				t.Errorf("InitRepo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
