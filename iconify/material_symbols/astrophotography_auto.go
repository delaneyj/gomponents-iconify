package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AstrophotographyAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 9l-1.25-2.75L15 5l2.75-1.25L19 1l1.25 2.75L23 5l-2.75 1.25L19 9Zm0 14l-1.25-2.75L15 19l2.75-1.25L19 15l1.25 2.75L23 19l-2.75 1.25L19 23ZM4.8 16h1.9l.7-2h3.2l.7 2h1.9L10 7H8l-3.2 9Zm3.05-3.35L9 9l1.15 3.65h-2.3ZM9 20q-3.35 0-5.675-2.325T1 12q0-3.35 2.325-5.675T9 4q3.35 0 5.675 2.325T17 12q0 3.35-2.325 5.675T9 20Z"/>`),
		g.Group(children),
	)
}