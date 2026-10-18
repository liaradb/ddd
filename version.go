package ddd

type Version int64

func (v Version) Value() int64       { return int64(v) }
func (v Version) Increment() Version { return Version(v + 1) }
