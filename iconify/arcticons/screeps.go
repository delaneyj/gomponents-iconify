package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screeps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 35.48L14.78 24L4.5 12.51h7.09L21.88 24L11.59 35.48Zm35.51 0H25.13a3.2 3.2 0 0 1 0-6.38H40a3.2 3.2 0 1 1 .57 6.38a5.42 5.42 0 0 1-.57 0Z"/>`),
		g.Group(children),
	)
}