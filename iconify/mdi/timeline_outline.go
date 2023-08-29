package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimelineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v6H2V2h2M2 22v-6h2v6H2m3-10c0 1.11-.89 2-2 2a2 2 0 1 1 2-2m19-6v12c0 1.11-.89 2-2 2H10a2 2 0 0 1-2-2v-4l-2-2l2-2V6a2 2 0 0 1 2-2h12c1.11 0 2 .89 2 2M10 6v12h12V6H10Z"/>`),
		g.Group(children),
	)
}