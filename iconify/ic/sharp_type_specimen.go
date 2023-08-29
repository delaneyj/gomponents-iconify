package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTypeSpecimen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6H2v16h16v-2H4z"/><path fill="currentColor" d="M22 2H6v16h16V2zm-5.37 12.5l-.8-2.3H12.2l-.82 2.3H9.81l3.38-9h1.61l3.38 9h-1.55z"/><path fill="currentColor" d="m13.96 7.17l-1.31 3.72h2.69l-1.3-3.72z"/>`),
		g.Group(children),
	)
}