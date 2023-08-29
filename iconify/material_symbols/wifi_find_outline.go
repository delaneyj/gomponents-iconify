package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFindOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 8.975q2.325-2.325 5.388-3.65T12 4q3.55 0 6.613 1.325T24 8.975L22.6 10.4q-2.025-2.025-4.75-3.213T12 6q-2.575 0-4.875.813t-4.2 2.262l10.475 10.5L12 21Zm9.6-1l-2.575-2.55q-.45.275-.95.413T17 18q-1.7 0-2.85-1.15T13 14q0-1.7 1.15-2.85T17 10q1.7 0 2.85 1.15T21 14q0 .575-.137 1.075t-.413.95L23 18.6L21.6 20ZM17 16q.85 0 1.425-.575T19 14q0-.85-.575-1.425T17 12q-.85 0-1.425.575T15 14q0 .85.575 1.425T17 16Zm-5 3.575Z"/>`),
		g.Group(children),
	)
}