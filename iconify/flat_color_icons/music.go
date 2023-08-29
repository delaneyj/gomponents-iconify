package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#E91E63"><circle cx="19" cy="33" r="9"/><path d="M24 6v27h4V14l11 3v-7z"/></g>`),
		g.Group(children),
	)
}