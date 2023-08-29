package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Load(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 1h1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12v-1h2v1a9 9 0 1 0 9-9h-1V1Z"/>`),
		g.Group(children),
	)
}