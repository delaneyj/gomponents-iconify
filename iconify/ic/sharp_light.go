package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6.06V3h-2v3.06A9.006 9.006 0 0 0 3.22 17H8c0 2.21 1.79 4 4 4s4-1.79 4-4h4.78A9.006 9.006 0 0 0 13 6.06zM12 15H5c0-3.86 3.14-7 7-7s7 3.14 7 7h-7z"/>`),
		g.Group(children),
	)
}