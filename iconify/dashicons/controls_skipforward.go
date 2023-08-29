package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlsSkipforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 12.4L2 16V4l6 3.6V4l8 4.8V4h2v12h-2v-4.8L8 16v-3.6z"/>`),
		g.Group(children),
	)
}