package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandLeetcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13h7.5M9.424 7.268l4.999-4.999m2.21 14.375l-2.402 2.415a3.189 3.189 0 0 1-4.524 0l-3.77-3.787a3.223 3.223 0 0 1 0-4.544l3.77-3.787a3.189 3.189 0 0 1 4.524 0l2.302 2.313"/>`),
		g.Group(children),
	)
}