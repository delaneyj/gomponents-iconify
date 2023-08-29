package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1 1l5 6l4.94-6H1zm4.94 9L1 16h10zm-2.82 5l2.83-3.44l3 3.44H3.12zM10 8h1v1h-1V8zm2 0h1v1h-1V8zM8 8h1v1H8V8zM6 8h1v1H6V8zM4 8h1v1H4V8zM2 8h1v1H2V8zM0 8h1v1H0V8z"/><path fill="currentColor" d="M15 8.47a4.807 4.807 0 0 1-1.879 3.632L12 11v3h3l-1.18-1.18A5.757 5.757 0 0 0 16 8.478V8.47a6.062 6.062 0 0 0-2.884-4.905l-.596.805a5.095 5.095 0 0 1 2.479 4.087z"/>`),
		g.Group(children),
	)
}