package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.966 6.453c.038-.248.238-.453.489-.453c.3 0 .549.246.508.544A4 4 0 1 1 9 3.354V2.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h.736a3 3 0 1 0 .73 2.453Z"/>`),
		g.Group(children),
	)
}