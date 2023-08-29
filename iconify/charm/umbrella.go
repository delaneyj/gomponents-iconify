package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25s.5-6.5 6.25-6.5s6.25 6.5 6.25 6.5zm6 .5v4s0 1.5 1.5 1.5s1.5-1.5 1.5-1.5"/>`),
		g.Group(children),
	)
}