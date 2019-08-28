package log

import "testing"

func TestInit(t *testing.T) {

}

func Test_output_isFile(t *testing.T) {
	tests := []struct {
		name string
		o    output
		want bool
	}{
		{
			o: "file",
			want: true,
		},
		{
			o: "stdout",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.isFile(); got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}