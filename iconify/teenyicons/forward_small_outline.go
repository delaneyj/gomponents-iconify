package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M8.5 9.5v-4l3.5 2l-3.5 2Zm-4 0v-4l3.5 2l-3.5 2Z"/>`),
		g.Group(children),
	)
}