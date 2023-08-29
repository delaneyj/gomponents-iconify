package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutPanelLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1 2l1-1h12l1 1v12l-1 1H2l-1-1V2Zm1 0v8h8V2H2Zm9 0v12h3V2h-3Z"/>`),
		g.Group(children),
	)
}