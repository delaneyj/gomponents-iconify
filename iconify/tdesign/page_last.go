package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageLast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 5v14h-2V5h2Zm-9.75.336L14.414 12L7.75 18.664L6.336 17.25l5.25-5.25l-5.25-5.25L7.75 5.336Z"/>`),
		g.Group(children),
	)
}