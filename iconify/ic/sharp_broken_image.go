package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBrokenImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3v8.59l-3-3.01l-4 4.01l-4-4l-4 4l-3-3.01V3h18zm-3 8.42l3 3.01V21H3v-8.58l3 2.99l4-4l4 4l4-3.99z"/>`),
		g.Group(children),
	)
}