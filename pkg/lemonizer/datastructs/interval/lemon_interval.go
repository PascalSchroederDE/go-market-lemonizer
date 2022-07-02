package lemon_interval

type LemonInterval struct{ *string }

func (e LemonInterval) String() string {
	if e.string == nil {
		return "<void>"
	}
	return *e.string
}

var (
	es = []string{"m1", "h1", "d1"}

	Invalid = LemonInterval{}
	Min     = LemonInterval{&es[0]}
	Hour    = LemonInterval{&es[1]}
	Day     = LemonInterval{&es[2]}
)
