package golazylog

type Color string

const (
	COLOR_B_BLACK  Color = "\x1b[30;1m"
	COLOR_B_RED    Color = "\x1b[31;1m"
	COLOR_B_GREEN  Color = "\x1b[32;1m"
	COLOR_B_YELLOW Color = "\x1b[33;1m"
	COLOR_B_BLUE   Color = "\x1b[34;1m"
	COLOR_B_MAROON Color = "\x1b[35;1m"
	COLOR_B_CYAN   Color = "\x1b[36;1m"
	COLOR_B_WHITE  Color = "\x1b[37;1m"

	COLOR_BLACK  Color = "\x1b[30m"
	COLOR_RER    Color = "\x1b[31m"
	COLOR_GREEN  Color = "\x1b[32m"
	COLOR_YELLOW Color = "\x1b[33m"
	COLOR_BLUE   Color = "\x1b[34m"
	COLOR_MAROON Color = "\x1b[35m"
	COLOR_CYAN   Color = "\x1b[36m"
	COLOR_WHITE  Color = "\x1b[37m"

	COLOR_NONE Color = "\x1b[0m"
)

func (cc ColorCode) String() string {
	return string(cc)
}

func Color(c ColorCode, s string) string {
	return c.String() + s
}

func ColorClear(c ColorCode, s string) string {
	return c.String() + s + COLOR_NONE.String()
}
