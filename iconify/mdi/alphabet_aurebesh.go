package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabetAurebesh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4v7h11.23L22 4h-3l-5.54 5H5V4H3m0 9v7h2v-5h8.46L19 20h3l-7.77-7H3Z"/>`),
		g.Group(children),
	)
}