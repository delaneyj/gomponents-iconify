package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l16 8.05v22.9l-16 8l-16-8V28l16 8l9.42-4.75v-8L24 28L8 20v-7.45Zm0 7.42l-8.62 4.34L24 20.61l8.62-4.35Z"/>`),
		g.Group(children),
	)
}