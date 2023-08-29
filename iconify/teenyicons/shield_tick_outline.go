package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldTickOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M4 7.5L7 10l4-5M7.5.5l-7 4v.72a9.651 9.651 0 0 0 7 9.28a9.651 9.651 0 0 0 7-9.28V4.5l-7-4Z"/>`),
		g.Group(children),
	)
}