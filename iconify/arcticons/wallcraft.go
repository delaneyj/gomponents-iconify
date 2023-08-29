package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.62 34.75L24 45.5L5.38 34.75v-21.5L24 2.5l18.62 10.75v21.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.5 16l-4.75 16L24 16l-4.75 16l-4.75-16"/>`),
		g.Group(children),
	)
}