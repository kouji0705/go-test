package sample

import (
	"fmt"
	"testing"

	pr "github.com/kouji0705/go-test/product"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		expected int
	}{
		{
			name: "1+1(成功)",
			args: args{1, 1},
			expected: 2,
		},
		{
			name: "1+1(失敗)",
			args: args{1, 1},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := pr.Add(tt.args.a, tt.args.b); actual != tt.expected {
				fmt.Println("テスト失敗", tt.name)
				t.Errorf("実行値と期待値が異なります。実行値: %v, 期待値: %v", actual, tt.expected)
			}else {
				fmt.Println("テスト成功", tt.name)
			}
		})
	}
}