package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePowerpointSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.6l-.3-.3l-6-6l-.3-.3H6zm2 2h10v6h6v16H8V5zm12 1.4L22.6 9H20V6.4zM13 13v2h4c1.2 0 2 .8 2 2s-.8 2-2 2s-2-.8-2-2h-2v7h2v-3.6c.6.4 1.3.6 2 .6c2.2 0 4-1.8 4-4s-1.8-4-4-4h-4z"/>`),
		g.Group(children),
	)
}