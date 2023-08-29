package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.5 17H4a4 4 0 0 0 2-3v-3a7 7 0 0 1 4-6a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3c.047.386.149.758.3 1.107M9 17v1a3 3 0 0 0 3.504 2.958M16 19h6"/>`),
		g.Group(children),
	)
}