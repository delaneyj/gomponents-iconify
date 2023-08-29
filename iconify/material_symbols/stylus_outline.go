package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.175 21q-.525.125-.913-.263T3 19.826l1-4.775L8.95 20l-4.775 1Zm4.775-1L4 15.05L15.45 3.6q.575-.575 1.425-.575T18.3 3.6l2.1 2.1q.575.575.575 1.425T20.4 8.55L8.95 20Zm7.925-15L6.525 15.35l2.125 2.125L19 7.125L16.875 5Z"/>`),
		g.Group(children),
	)
}