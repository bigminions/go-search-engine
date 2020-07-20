package analyzer

import (
	"reflect"
	"sort"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{doc: "hello world, how are you!"},
			want: []string{"hello", "world", "how", "are", "you"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_analyze(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{doc: "hello world, how are you!"},
			want: []string{"hello", "world"},
		},
		{
			name: "test2",
			args: args{doc: ""},
			want: []string{},
		},
		{
			name: "test3",
			args: args{doc: " “Anatomy Park” was a great premise for an episode, as Rick shrinks down Morty and sends him into the tiny theme park he’s built inside a homeless man’s body. "},
			want: []string{"anatomy","great","premise","episode","rick","shrinks","morty","sends","tiny","theme","park","built","inside","homeless","man","body"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyze(tt.args.doc); !equal(got, tt.want) {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(got []string, want []string) bool {
	sort.Strings(got)
	sort.Strings(want)
	return reflect.DeepEqual(got, want)
}