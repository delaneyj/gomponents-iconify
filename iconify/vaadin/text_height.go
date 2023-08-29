package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 3h1l-1.5-3L13 3h1v10h-1l1.5 3l1.5-3h-1zM1 0v3h4v13h3V3h4V0z"/>`),
		g.Group(children),
	)
}