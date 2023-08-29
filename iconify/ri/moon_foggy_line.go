package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonFoggyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20.335v-2.2a7.523 7.523 0 0 0 3.623-4.281a9 9 0 0 1-10.622-8.99A7.518 7.518 0 0 0 5.151 10H3.117a9.505 9.505 0 0 1 8.538-7.963a7 7 0 0 0 10.316 8.728A9.503 9.503 0 0 1 16 20.335ZM7 20h7v2H7v-2Zm-3-8h6v2H4v-2Zm-2 4h10v2H2v-2Z"/>`),
		g.Group(children),
	)
}