package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointerTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13 10V8.5A1.5 1.5 0 0 0 11.5 7v0A1.5 1.5 0 0 0 10 8.5V11m3 0V9.5A1.5 1.5 0 0 1 14.5 8v0A1.5 1.5 0 0 1 16 9.5v.5m-9 1.5v-7A1.5 1.5 0 0 1 8.5 3v0A1.5 1.5 0 0 1 10 4.5V11m0 1v-1m-3 3v-2.5A1.5 1.5 0 0 0 5.5 10v0A1.5 1.5 0 0 0 4 11.5v2C4 16 5.5 21 11.5 21c2.5 0 7.5-1.5 7.5-7.5V12a2 2 0 0 0-2-2h-1m0 0v2"/>`),
		g.Group(children),
	)
}