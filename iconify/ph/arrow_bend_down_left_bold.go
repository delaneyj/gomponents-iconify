package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDownLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 56a108.12 108.12 0 0 1-108 108H61l27.52 27.51a12 12 0 0 1-17 17l-48-48a12 12 0 0 1 0-17l48-48a12 12 0 0 1 17 17L61 140h67a84.09 84.09 0 0 0 84-84a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}