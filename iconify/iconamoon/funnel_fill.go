package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 0-.8 1.6l5.6 7.467a1 1 0 0 1 .2.6V20a1 1 0 0 0 1.447.894l4-2A1 1 0 0 0 15 18v-5.333a1 1 0 0 1 .2-.6L20.8 4.6A1 1 0 0 0 20 3H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}