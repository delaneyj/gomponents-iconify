package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pipeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9v1H5V4h1V2H1v2h1v7a2 2 0 0 0 2 2h5v1h2V9Zm13 11v-8a2 2 0 0 0-2-2h-5V9h-2v5h2v-1h4v7h-1v2h5v-2Z"/>`),
		g.Group(children),
	)
}