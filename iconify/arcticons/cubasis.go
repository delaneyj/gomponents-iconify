package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cubasis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.082" cy="24" r="5.232" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.811 23.81c0-6.468 5.042-11.795 11.415-12.27l-7.04-7.04L5.687 24l19.5 19.5l7.42-7.42c-6.563-.285-11.795-5.612-11.795-12.27Z"/>`),
		g.Group(children),
	)
}