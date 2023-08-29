package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m0 0l6 4l8-4l6 4v16l-6-4l-8 4l-6-4V0zm7 6v11l6-3V3L7 6z"/>`),
		g.Group(children),
	)
}