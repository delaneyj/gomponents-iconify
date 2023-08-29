package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorArrowClockwiseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5.25A3.25 3.25 0 0 1 7.25 2h9.5A3.25 3.25 0 0 1 20 5.25v5.025a1.756 1.756 0 0 0-.232.294A7 7 0 0 0 12.101 22H7.25A3.25 3.25 0 0 1 4 18.751V5.25ZM9 5a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H9Zm.5 8.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0ZM8.25 18.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm13-3.5a.75.75 0 0 0 .75-.75V11.5a.75.75 0 0 0-1.5 0v.626A6 6 0 1 0 23 17a.75.75 0 0 0-1.5 0a4.5 4.5 0 1 1-1.688-3.513l.008.006a.354.354 0 0 0 .01.007H18.5a.75.75 0 0 0 0 1.5h2.75Z"/>`),
		g.Group(children),
	)
}