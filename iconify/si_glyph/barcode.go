package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 2h1.953v10.914H14zm-3 0h1.967v10.914H11zM9 2h.95v10.914H9zM5 2h1.972v10.914H5zM3 2h.973v10.914H3zM1 2h.973v10.914H1z"/>`),
		g.Group(children),
	)
}