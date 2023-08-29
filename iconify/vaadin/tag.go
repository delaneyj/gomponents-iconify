package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 1H1v7l7 7l7-7zM3.75 5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5z"/>`),
		g.Group(children),
	)
}