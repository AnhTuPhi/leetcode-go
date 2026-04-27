package runner

// TestCase giữ input, expected output và tên tuỳ chọn cho mỗi test.
type TestCase[I, O any] struct {
	Name     string
	Input    I
	Expected O
}

// Case tạo TestCase có tên — dùng khi muốn label rõ ràng.
func Case[I, O any](name string, input I, expected O) TestCase[I, O] {
	return TestCase[I, O]{Name: name, Input: input, Expected: expected}
}

// Of tạo TestCase không tên — dùng khi test case đơn giản.
func Of[I, O any](input I, expected O) TestCase[I, O] {
	return TestCase[I, O]{Input: input, Expected: expected}
}
