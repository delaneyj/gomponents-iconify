package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19 10l-4 4v-3h-4v4h3l-4 4l-4-4h3v-4H5v3l-4-4l4-4v3h4V5H6l4-4l4 4h-3v4h4V6z"/>`),
		g.Group(children),
	)
}