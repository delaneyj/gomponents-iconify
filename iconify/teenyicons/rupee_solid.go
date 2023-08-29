package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1V0h11v1H8.329a4.494 4.494 0 0 1 1.644 3H13v1H9.973A4.5 4.5 0 0 1 5.5 9H3.852l5.973 5.12l-.65.76l-7-6A.5.5 0 0 1 2.5 8h3a3.5 3.5 0 0 0 3.465-3H2V4h6.965A3.5 3.5 0 0 0 5.5 1H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}