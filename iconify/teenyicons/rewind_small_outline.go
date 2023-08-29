package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.5 9.5v-4L3 7.5l3.5 2Zm4 0v-4L7 7.5l3.5 2Z"/>`),
		g.Group(children),
	)
}