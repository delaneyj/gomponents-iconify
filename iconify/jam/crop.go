package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17H7a4 4 0 0 1-4-4V5H1a1 1 0 1 1 0-2h2v-4v2a1 1 0 1 1 2 0v2h8a4 4 0 0 1 4 4v8h2a1 1 0 0 1 0 2h-2v4v-2a1 1 0 0 1-2 0v-2zm0-2V7a2 2 0 0 0-2-2H5v8a2 2 0 0 0 2 2h8z"/>`),
		g.Group(children),
	)
}