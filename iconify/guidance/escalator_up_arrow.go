package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m10.5.5l-10 10m10-10c-.398.398-1.133.654-1.79.812c-.878.212-1.79.25-2.687.149c-.697-.08-1.464-.235-1.773-.544M10.5.5c-.398.398-.654 1.133-.812 1.79a7.745 7.745 0 0 0-.149 2.687c.08.697.235 1.464.544 1.773M5.515 23.5H.5v-6H6l8.243-8.243A6 6 0 0 1 18.485 7.5H23.5v6H18l-8.243 8.243A6 6 0 0 1 5.515 23.5Z"/>`),
		g.Group(children),
	)
}