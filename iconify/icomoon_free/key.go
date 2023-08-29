package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 0a5 5 0 0 0-4.916 5.916L0 12v3a1 1 0 0 0 1 1h1v-1h2v-2h2v-2h2l1.298-1.298A5 5 0 1 0 11 0zm1.498 5.002a1.5 1.5 0 1 1 .001-3.001a1.5 1.5 0 0 1-.001 3.001z"/>`),
		g.Group(children),
	)
}