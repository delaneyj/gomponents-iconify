package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutSidebarLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1H2Zm12 13H7V2h7v12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}