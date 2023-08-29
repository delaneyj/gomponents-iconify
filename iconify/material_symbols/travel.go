package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Travel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.85 17.15l-3.2-1.75l1.05-1.05l2.5.35l3.9-3.9l-7.8-4.25l1.4-1.4l9.55 2.45l3.925-3.875Q18.6 3.3 19.238 3.3t1.062.425q.425.425.425 1.063T20.3 5.85l-3.9 3.9l2.45 9.55l-1.4 1.4l-4.25-7.8l-3.9 3.9l.35 2.5l-1.05 1.05l-1.75-3.2Z"/>`),
		g.Group(children),
	)
}