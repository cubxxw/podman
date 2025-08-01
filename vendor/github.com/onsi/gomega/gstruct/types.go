package gstruct

// Options is the type for options passed to some matchers.
type Options int

const (
	//IgnoreExtras tells the matcher to ignore extra elements or fields, rather than triggering a failure.
	IgnoreExtras Options = 1 << iota
	//IgnoreMissing tells the matcher to ignore missing elements or fields, rather than triggering a failure.
	IgnoreMissing
	//AllowDuplicates tells the matcher to permit multiple members of the slice to produce the same ID when
	//considered by the identifier function. All members that map to a given key must still match successfully
	//with the matcher that is provided for that key.
	AllowDuplicates
	//IgnoreUnexportedExtras tells the matcher to ignore extra unexported fields, rather than triggering a failure.
	//it is not possible to check the value of unexported fields, so this option is only useful when you want to
	//check every exported fields, but you don't care about extra unexported fields.
	IgnoreUnexportedExtras
)
