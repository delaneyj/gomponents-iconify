package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 44c-.69 0-1.25-.56-1.25-1.25V10.304L10.14 23.126a1.25 1.25 0 1 1-1.782-1.752L23.097 6.386l.026-.027a1.245 1.245 0 0 1 1.009-.352a1.245 1.245 0 0 1 .785.393l14.724 14.974a1.25 1.25 0 1 1-1.782 1.752l-12.61-12.822V42.75c0 .69-.559 1.25-1.25 1.25Z"/>`),
		g.Group(children),
	)
}