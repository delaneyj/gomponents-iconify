package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessBishop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 16l-1.447.724a1 1 0 0 0-.553.894V20h12v-2.382a1 1 0 0 0-.553-.894L16 16H8zm3-12a1 1 0 1 0 2 0a1 1 0 1 0-2 0M9.5 16C7.833 16 7 14.331 7 13c0-3.667 1.667-6 5-7c3.333 1 5 3.427 5 7c0 1.284-.775 2.881-2.325 3H9.5zM15 8l-3 3m0-6v1"/>`),
		g.Group(children),
	)
}