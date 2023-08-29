package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1h16v3H0V1zm1 4v11h14V5H1zm10 4H5V7h6v2z"/>`),
		g.Group(children),
	)
}