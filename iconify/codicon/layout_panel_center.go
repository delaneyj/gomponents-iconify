package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutPanelCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1H2Zm0 13V2h2v12H2Zm3-4V2h6v8H5Zm7-8h2v12h-2V2Z"/>`),
		g.Group(children),
	)
}