package tesify

import (
	"golang_edu/third-party-package/tesify/mocks"
	"testing"
)

func TestFoo(t *testing.T) {
	type args struct {
		adder Adder
		x     int
		y     int
	}

	sanityAdder := new(mocks.Adder)
	sanityAdder.On("Add", 10, 11).Return(22)

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"sanity",
			args{sanityAdder, 10, 11},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.args.adder, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
