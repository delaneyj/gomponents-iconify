package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stylus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.2 20.975q-.525.125-.913-.263t-.262-.912l.875-4.25l4.55 4.55l-4.25.875Zm5.875-2.1l-4.95-4.95L15.45 3.6q.575-.575 1.425-.575T18.3 3.6l2.1 2.1q.575.575.575 1.425T20.4 8.55L10.075 18.875Z"/>`),
		g.Group(children),
	)
}