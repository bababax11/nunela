package nunela_test

import (
	"testing"

	"github.com/bababax11/nunela"
	"github.com/vorduin/nune"
)

func TestTensorDotWithOneAxis(t *testing.T) {
	nune.EnvConfig.Interactive = true
	cases := map[string]struct {
		tensor0  nune.Tensor[int]
		tensor1  nune.Tensor[int]
		axes     []int
		expected nune.Tensor[int]
	}{
		"matrix multiplation": {
			nune.Range[int](0, 8, 1).Reshape(2, 4),
			nune.Range[int](0, 12, 1).Reshape(4, 3),
			[]int{1, 0},
			nune.FromBuffer([]int{
				42, 48, 54,
				114, 136, 158,
			}).Reshape(2, 3),
		},
		"3D tensordot with different axes": {
			nune.Range[int](0, 8, 1).Reshape(2, 4),
			nune.Range[int](0, 12, 1).Reshape(2, 2, 3),
			[]int{0, 1},
			nune.FromBuffer([]int{
				12, 16, 20,
				36, 40, 44,

				15, 21, 27,
				51, 57, 63,

				18, 26, 34,
				66, 74, 82,

				21, 31, 41,
				81, 91, 101,
			}).Reshape(4, 2, 3),
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			val := nunela.TensorDotWithOneAxis([]*nune.Tensor[int]{
				&tt.tensor0, &tt.tensor1,
			}, tt.axes)
			if !nunela.Equal(val, &tt.expected) {
				t.Error(name, val.Ravel(), tt.expected.Ravel())
			}
		})
	}
}
