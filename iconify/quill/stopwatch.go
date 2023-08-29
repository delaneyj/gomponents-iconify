package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.9 7.869a15 15 0 0 0-2.769-2.77m-7.434-.988a13 13 0 0 0-3.394 0m12.304 2.282l-2.829 2.829M16 17l5 5m6-5c0 6.075-4.925 11-11 11S5 23.075 5 17S9.925 6 16 6s11 4.925 11 11Z"/>`),
		g.Group(children),
	)
}