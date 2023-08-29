package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRotateRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20a2 2 0 0 1-2-2V9c0-1.1.9-2 2-2h4l4 4v7c0 1.11-.89 2-2 2H5m6-8.17L8.17 9H8v3h3v-.17m9-.83h3l-4 4l-4-4h3c0-3.31-2.69-6-6-6l-1 .08V3.06L12 3c4.42 0 8 3.58 8 8"/>`),
		g.Group(children),
	)
}