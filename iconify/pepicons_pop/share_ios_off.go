package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareIosOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M9 3a1 1 0 0 1 2 0v10.5a1 1 0 1 1-2 0V3Z"/><path d="M10.64 2.768a1 1 0 1 0-1.28-1.536l-3 2.5a1 1 0 0 0 1.28 1.536l3-2.5Z"/><path d="M9.36 2.768a1 1 0 1 1 1.28-1.536l3 2.5a1 1 0 1 1-1.28 1.536l-3-2.5ZM13 9a1 1 0 1 1 0-2h1c1.623 0 3 1.165 3 2.692v7.616C17 18.835 15.623 20 14 20H6c-1.623 0-3-1.165-3-2.692V9.692C3 8.165 4.377 7 6 7h1a1 1 0 0 1 0 2H6c-.586 0-1 .35-1 .692v7.616c0 .342.414.692 1 .692h8c.586 0 1-.35 1-.692V9.692C15 9.35 14.586 9 14 9h-1Z"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}