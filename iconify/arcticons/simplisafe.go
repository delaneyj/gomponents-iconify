package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplisafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.21 16.71a12.21 12.21 0 0 0-24.42 0v5h24.42v9.58a12.21 12.21 0 0 1-24.42 0v-2.54"/>`),
		g.Group(children),
	)
}