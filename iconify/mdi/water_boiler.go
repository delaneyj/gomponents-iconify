package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBoiler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2c-1.11 0-2 .89-2 2v12c0 1.11.89 2 2 2h1v2H6v2h3c1.11 0 2-.89 2-2v-2h2v2c0 1.11.89 2 2 2h3v-2h-3v-2h1c1.11 0 2-.89 2-2V4c0-1.11-.89-2-2-2H8m4 2.97a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2m-2 9.53h4V16h-4v-1.5Z"/>`),
		g.Group(children),
	)
}