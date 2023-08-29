package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gsmarena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.823 22.28L21.591 4.5h16.705v27.947l-10.373-3.839l-.261-14.209l-8.094 11.118Zm5.037 11.053L38.297 43.5H9.704Z"/>`),
		g.Group(children),
	)
}