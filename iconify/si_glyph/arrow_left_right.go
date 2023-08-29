package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 9h8.066v2.864l3.85-3.82l-3.85-4.013L13 7H5l-.114-2.97L1.03 8.07l3.856 3.866L5 9z"/>`),
		g.Group(children),
	)
}