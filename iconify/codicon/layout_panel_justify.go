package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutPanelJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1H2Zm0 9V2h2v8H2Zm3 0V2h6v8H5Zm7 0V2h2v8h-2Z"/>`),
		g.Group(children),
	)
}