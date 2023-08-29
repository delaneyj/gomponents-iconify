package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6A1.5 1.5 0 0 1 12 1.5V5H3V1.5ZM1.5 6A1.5 1.5 0 0 0 0 7.5v4A1.5 1.5 0 0 0 1.5 13H2V9h11v4h.5a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 6h-12Z"/><path fill="currentColor" d="M3 10h9v5H3v-5Z"/>`),
		g.Group(children),
	)
}