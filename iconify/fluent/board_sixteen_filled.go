package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 2v3H2v-.5A2.5 2.5 0 0 1 4.5 2h3Zm1 0v8H14V4.5A2.5 2.5 0 0 0 11.5 2h-3Zm5.5 9H8.5v3h3a2.5 2.5 0 0 0 2.5-2.5V11Zm-6.5 3V6H2v5.5A2.5 2.5 0 0 0 4.5 14h3Z"/>`),
		g.Group(children),
	)
}