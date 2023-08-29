package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.45 19h9.1l3.075-9.225L12 4.45L4.375 9.775L7.45 19ZM6 21L2 9l10-7l10 7l-4 12H6Zm6-9.275Z"/>`),
		g.Group(children),
	)
}