package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Melvoridle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8.5a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h9.56V26.307l8.21 8.218l8.21-8.218V39.5h9.02a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2H30.873L24 16.404L17.127 8.5Zm17.36 7.9l.28 17.317"/>`),
		g.Group(children),
	)
}