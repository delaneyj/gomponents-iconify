package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v4.384l-2.43 1.223a1 1 0 1 0 .898 1.786l2.981-1.5A.999.999 0 0 0 13 12V7a1 1 0 0 0-1-1Zm0-4a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}