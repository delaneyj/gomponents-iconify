package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Story(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 5a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2h-7a2 2 0 0 1-2-2h9z"/><path fill="currentColor" d="M13 4H5a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2z"/><path fill="currentColor" d="M7 16h8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z"/>`),
		g.Group(children),
	)
}