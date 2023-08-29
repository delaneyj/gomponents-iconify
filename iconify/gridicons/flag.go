package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6a2 2 0 0 0-2-2H5v17h2v-7h5a2 2 0 0 0 2 2h6V6h-5z"/>`),
		g.Group(children),
	)
}