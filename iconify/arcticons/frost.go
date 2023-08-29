package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5L5.38 13.25v21.5L24 45.5l18.62-10.75v-21.5Zm-3.39 10.23h6.78v22.54h-6.78Z"/>`),
		g.Group(children),
	)
}