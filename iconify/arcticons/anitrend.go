package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anitrend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.012 42.5l-12.024-37L5.5 42.5m4.162-12.488H25.85M17.988 5.5H42.5m-12.025 37v-37"/>`),
		g.Group(children),
	)
}