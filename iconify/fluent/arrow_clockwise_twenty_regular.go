package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.066 9.05a7 7 0 0 1 12.557-3.22l.126.17H12.5a.5.5 0 1 0 0 1h4a.5.5 0 0 0 .5-.5V2.502a.5.5 0 0 0-1 0v2.207a8 8 0 1 0 1.986 4.775a.5.5 0 0 0-.998.064A7 7 0 1 1 3.066 9.05Z"/>`),
		g.Group(children),
	)
}