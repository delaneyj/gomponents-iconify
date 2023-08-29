package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandXdeep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.401 8.398L16 6h5l-4 6l4 6h-5L8 6H3l4 6l-4 6h5l1.596-2.393"/>`),
		g.Group(children),
	)
}