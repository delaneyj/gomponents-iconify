package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDashTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 3A3.5 3.5 0 0 0 3 6.5v1.864a1 1 0 0 0 2 0V6.5A1.5 1.5 0 0 1 6.5 5h1.864a1 1 0 0 0 0-2H6.5Zm9.136 0a1 1 0 1 0 0 2H17.5A1.5 1.5 0 0 1 19 6.5v1.864a1 1 0 1 0 2 0V6.5A3.5 3.5 0 0 0 17.5 3h-1.864ZM5 15.636a1 1 0 1 0-2 0V17.5A3.5 3.5 0 0 0 6.5 21h1.864a1 1 0 1 0 0-2H6.5A1.5 1.5 0 0 1 5 17.5v-1.864Zm16 0a1 1 0 1 0-2 0V17.5a1.5 1.5 0 0 1-1.5 1.5h-1.864a1 1 0 1 0 0 2H17.5a3.5 3.5 0 0 0 3.5-3.5v-1.864ZM8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8Z"/>`),
		g.Group(children),
	)
}