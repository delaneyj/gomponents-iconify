package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.25 19.45q-.575-.225-.913-.725T4 17.6V4q0-1.05.875-1.638T6.75 2.15l5.95 2.375q.575.225.938.738T14 6.4V20q0 1.05-.875 1.65t-1.875.2l-6-2.4ZM16 19V6.4q0-1.675-1.35-2.713T11.675 2H18q.825 0 1.413.588T20 4v13q0 .825-.588 1.413T18 19h-2Z"/>`),
		g.Group(children),
	)
}