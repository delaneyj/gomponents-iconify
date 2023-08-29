package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 2h14v2L9 9v7l-2-2V9L1 4V2zm0-2h14v1H1V0z"/>`),
		g.Group(children),
	)
}