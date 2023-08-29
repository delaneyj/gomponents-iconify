package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shoutem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.849 6.13l18.026 18.453a2.092 2.092 0 0 1-.657 3.378L8.468 42.317a2.092 2.092 0 0 1-2.748-2.772l14.724-32.81a2.092 2.092 0 0 1 3.405-.605Z"/>`),
		g.Group(children),
	)
}