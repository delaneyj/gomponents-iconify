package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lambda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 20l4.16-12.09L9.34 6H8V4h2c.42 0 .78.26.93.63L16.66 18H18v2h-2c-.43 0-.79-.27-.93-.64l-3.74-8.71L8.12 20H6Z"/>`),
		g.Group(children),
	)
}