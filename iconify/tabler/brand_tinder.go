package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.918 8.174c2.56 4.982.501 11.656-5.38 12.626C5.836 22.487.698 13.084 6.484 7.571C6.793 7.266 7.645 6.476 8 6.222c0 .528.27 3.475 1 3.167c3 0 4-4.222 3.587-7.389c2.7 1.411 4.987 3.376 6.331 6.174z"/>`),
		g.Group(children),
	)
}