package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 76H40a12 12 0 0 0-12 12v112a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V88a12 12 0 0 0-12-12Zm4 124a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4V88a4 4 0 0 1 4-4h144a4 4 0 0 1 4 4Zm40-144v120a4 4 0 0 1-8 0V56a4 4 0 0 0-4-4H64a4 4 0 0 1 0-8h152a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}