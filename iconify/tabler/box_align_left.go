package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.002 20.003v-16h-5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h5zm5 0h-.01m5.011 0h-.011m.011-5.001h-.011m.011-6h-.011m.011-5h-.011m-4.99 0h-.01"/>`),
		g.Group(children),
	)
}