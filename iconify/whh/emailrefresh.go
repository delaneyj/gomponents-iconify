package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emailrefresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m667 640l350-349q7 14 7 29v640q0 15-7 29zM63 256h135q23-111 111-183.5T512 0t203 72.5T826 256h135L512 705zm449-128q-45 0-84 20t-66 54l-10-10q-13 0-22.5 9.5T320 224v64q0 13 9.5 22.5T352 320h64q13 0 22.5-9.5T448 288l-41-41q18-25 45.5-40t59.5-15q37 0 67.5 19.5T626 263l57-29q-24-48-70-77t-101-29zm150 310l10 10q13 0 22.5-9.5T704 416v-64q0-13-9.5-22.5T672 320h-64q-13 0-22.5 9.5T576 352l41 41q-39 55-105 55q-37 0-67.5-19.5T398 377l-57 29q24 48 70 77t101 29q45 0 84-20t66-54zM7 989q-7-14-7-29V320q0-15 7-29l350 349zm395-304l71 71q6 6 15.5 9t16.5 3h7q26 1 39-12l71-71l339 339H63z"/>`),
		g.Group(children),
	)
}