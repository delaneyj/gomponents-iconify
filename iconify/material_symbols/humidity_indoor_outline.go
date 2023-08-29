package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumidityIndoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 16.5q1.65 0 2.813-1.15T16 12.55q0-.8-.3-1.513t-.875-1.262L12 7L9.175 9.775q-.575.55-.875 1.263T8 12.55q0 1.65 1.188 2.8t2.837 1.15ZM10 12.5q0-.375.15-.7t.425-.6L12 9.8l1.425 1.4q.275.275.425.6t.15.7h-4ZM4 20V8l8-6l8 6v12H4Zm2-2h12V9l-6-4.5L6 9v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}