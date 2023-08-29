package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 19H6a5 5 0 0 1-.332-9.99a7.018 7.018 0 0 1 1.333-1.909M19 19L5 5m14 14l2 2M10 5.29A7.002 7.002 0 0 1 18.93 11H19a3.999 3.999 0 0 1 3.122 6.5"/>`),
		g.Group(children),
	)
}