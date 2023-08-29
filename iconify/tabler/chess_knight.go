package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessKnight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 16l-1.447.724a1 1 0 0 0-.553.894V20h12v-2.382a1 1 0 0 0-.553-.894L16 16H8zM9 3l1 3l-3.491 2.148A1 1 0 0 0 7.033 10H10l-2.073 6h7.961L16 11c0-3-1.09-5.983-4-7c-1.94-.678-2.94-1.011-3-1z"/>`),
		g.Group(children),
	)
}