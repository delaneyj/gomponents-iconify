package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shadow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.477 21.477 0 0 0 3.583 30.712a10.22 10.22 0 1 1 13.705 13.705A21.496 21.496 0 1 0 24 2.5Z"/>`),
		g.Group(children),
	)
}