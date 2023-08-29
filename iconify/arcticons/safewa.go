package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safewa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.926 33.691V11.69L30.764 7.5l-7.284 9.428l-10.406 4.19l1.04 7.334l2.602 6.81l-1.561 3.666l6.243 1.571l.446-1.047l13.082-5.761v0ZM36.5 5.5h4c1.108 0 2 .892 2 2v4m0 25v4c0 1.108-.892 2-2 2h-4m-25 0h-4c-1.108 0-2-.892-2-2v-4m0-25v-4c0-1.108.892-2 2-2h4"/>`),
		g.Group(children),
	)
}