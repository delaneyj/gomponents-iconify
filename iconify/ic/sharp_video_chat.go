package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpVideoChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20l4-4h16V2H2zm15 11l-2-1.99V14H7V6h8v2.99L17 7v6z"/>`),
		g.Group(children),
	)
}