package solutions

// PerDayInput is an interface that exposes a simple structure:
// The runner should call Prepare() on it to prepare the input
// Then call the part 1 and part 2 functions if wanted.
type PerDayInput interface {
	Prepare(input string)
	Part1() string
	Part2() string
}

var RegisteredDays = []struct{
	DummyInput PerDayInput
	StringInput *string
	ExpectedOutputs []string
}{

}