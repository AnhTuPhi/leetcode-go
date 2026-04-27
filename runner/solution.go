package runner

// Solution là interface chính — implement 2 method này cho mỗi bài.
type Solution[I, O any] interface {
	Solve(input I) O
	TestCases() []TestCase[I, O]
}

// Equaler là interface tuỳ chọn — implement khi cần custom equality.
// Ví dụ: []int cần sort trước khi so, [][]string không quan tâm thứ tự...
// Nếu Solution không implement Equaler, runner dùng reflect.DeepEqual.
type Equaler[O any] interface {
	AssertEqual(expected, actual O) bool
}
