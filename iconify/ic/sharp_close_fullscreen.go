package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCloseFullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 3.41L16.71 8.7L20 12h-8V4l3.29 3.29L20.59 2L22 3.41zM3.41 22l5.29-5.29L12 20v-8H4l3.29 3.29L2 20.59L3.41 22z"/>`),
		g.Group(children),
	)
}