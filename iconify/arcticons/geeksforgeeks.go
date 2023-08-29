package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geeksforgeeks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.994 24H43.5a9.751 9.751 0 1 1-2.857-6.894"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.006 24H4.5a9.751 9.751 0 1 0 2.857-6.894"/>`),
		g.Group(children),
	)
}