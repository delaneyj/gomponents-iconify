package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 5.5h5v-5h-10l5 5Zm0 0h-5v4l5 5v-4h5l-5-5Z"/>`),
		g.Group(children),
	)
}