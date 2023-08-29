package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 32a4 4 0 0 1-4 4H56a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4Zm-32 40v152a12 12 0 0 1-12 12H96a12 12 0 0 1-12-12V72a12 12 0 0 1 12-12h64a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4H96a4 4 0 0 0-4 4v152a4 4 0 0 0 4 4h64a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}