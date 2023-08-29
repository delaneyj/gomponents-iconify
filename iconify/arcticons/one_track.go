package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.221 26.715l5.21-2.865v19.211"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.304 34.249L24 44.5L5.696 34.249V13.751L24 3.5l18.304 10.251v20.498z"/>`),
		g.Group(children),
	)
}