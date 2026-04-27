package runner

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

// ANSI colors
const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

// Run chạy toàn bộ test case, in PASS/FAIL và thời gian mỗi case.
func Run[I, O any](sol Solution[I, O]) {
	cases := sol.TestCases()
	name := fmt.Sprintf("%T", sol)

	fmt.Printf("\n%s=== %s ===%s\n", bold+cyan, name, reset)

	passed := 0
	for i, tc := range cases {
		label := tc.Name
		if label == "" {
			label = fmt.Sprintf("Case %d", i+1)
		}

		inputSnapshot := fmt.Sprintf("%+v", tc.Input)
		start := time.Now()
		actual := sol.Solve(tc.Input)
		elapsed := time.Since(start)

		ok := assertEqual(sol, tc.Expected, actual)
		if ok {
			passed++
		}

		status := green + "PASS" + reset
		if !ok {
			status = red + "FAIL" + reset
		}

		fmt.Printf("  [%s] %-24s %s(%s)%s\n",
			status, label, yellow, fmtDuration(elapsed), reset)

		if !ok {
			fmt.Printf("       input:    %+v\n", inputSnapshot)
			fmt.Printf("       expected: %+v\n", tc.Expected)
			fmt.Printf("       actual:   %+v\n", actual)
		}
	}

	fmt.Printf("  %s%d/%d passed%s\n\n", bold, passed, len(cases), reset)
}

// Benchmark chạy warmup rồi đo iterations lần, in avg/min/max/p99.
func Benchmark[I, O any](sol Solution[I, O], warmup, iterations int) {
	cases := sol.TestCases()
	name := fmt.Sprintf("%T", sol)

	fmt.Printf("\n%s=== BENCHMARK: %s ===%s\n", bold+cyan, name, reset)
	fmt.Printf("  warmup=%d  iterations=%d\n\n", warmup, iterations)

	for i, tc := range cases {
		label := tc.Name
		if label == "" {
			label = fmt.Sprintf("Case %d", i+1)
		}

		// Warmup — bỏ qua kết quả để JIT ổn định
		for w := 0; w < warmup; w++ {
			sol.Solve(tc.Input)
		}

		times := make([]int64, iterations)
		for it := 0; it < iterations; it++ {
			start := time.Now()
			sol.Solve(tc.Input)
			times[it] = time.Since(start).Nanoseconds()
		}

		stats := computeStats(times)
		fmt.Printf("  %-24s  avg=%s  min=%s  max=%s  p99=%s\n",
			label,
			fmtNs(stats.avg),
			fmtNs(stats.min),
			fmtNs(stats.max),
			fmtNs(stats.p99),
		)
	}
	fmt.Println()
}

// --- helpers ---

// assertEqual dùng Equaler nếu solution implement, ngược lại dùng reflect.DeepEqual.
func assertEqual[I, O any](sol Solution[I, O], expected, actual O) bool {
	if e, ok := any(sol).(Equaler[O]); ok {
		return e.AssertEqual(expected, actual)
	}
	return reflect.DeepEqual(expected, actual)
}

type stats struct{ avg, min, max, p99 int64 }

func computeStats(times []int64) stats {
	sorted := make([]int64, len(times))
	copy(sorted, times)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	var sum int64
	for _, t := range times {
		sum += t
	}

	p99idx := int(math.Ceil(float64(len(sorted))*0.99)) - 1
	return stats{
		avg: sum / int64(len(times)),
		min: sorted[0],
		max: sorted[len(sorted)-1],
		p99: sorted[p99idx],
	}
}

func fmtDuration(d time.Duration) string {
	return fmtNs(d.Nanoseconds())
}

func fmtNs(ns int64) string {
	switch {
	case ns < 1_000:
		return fmt.Sprintf("%dns", ns)
	case ns < 1_000_000:
		return fmt.Sprintf("%.2fµs", float64(ns)/1_000)
	case ns < 1_000_000_000:
		return fmt.Sprintf("%.2fms", float64(ns)/1_000_000)
	default:
		return fmt.Sprintf("%.2fs", float64(ns)/1_000_000_000)
	}
}
