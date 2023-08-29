package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandKick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h5v4h3V6h2V4h6v4h-2v2h-2v4h2v2h2v4h-6v-2h-2v-2H9v4H4z"/>`),
		g.Group(children),
	)
}