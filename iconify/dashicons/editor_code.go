package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9 6l-4 4l4 4l-1 2l-6-6l6-6zm2 8l4-4l-4-4l1-2l6 6l-6 6z"/>`),
		g.Group(children),
	)
}