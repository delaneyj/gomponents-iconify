package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CustomPostType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zM6 6h5v5H6V6zm4.5 13a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5zm3-6l3-5l3 5h-6z"/>`),
		g.Group(children),
	)
}