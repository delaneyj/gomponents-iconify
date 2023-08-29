package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3.262 10.878l8 8.789c.4.44 1.091.44 1.491 0l8-8.79c.313-.344.349-.859.087-1.243L17.303 4.44a1 1 0 0 0-.823-.436H7.554a1 1 0 0 0-.823.436l-3.54 5.192c-.263.385-.227.901.087 1.246z"/>`),
		g.Group(children),
	)
}