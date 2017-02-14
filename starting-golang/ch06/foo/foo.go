// foo パッケージの概要説明。
package foo

const (
	// oo の最大値
	MAX = 100
)

const (
	A = iota
	B
	C
)

// T 型の定義
type T struct {
	Field1 int
	field2 string
}

// *T 型に定義された 1 番目のメソッド
func (t *T) Method1()  {
}

// *T 型に定義された 2 番目のメソッド
func (t *T) Method2()  {
}

// 型 コンストラクタ
func New() *T {
	return &T{}
}
