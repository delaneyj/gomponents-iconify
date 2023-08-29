package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6v2h2v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8h2V6Zm9 12H6v-2h5Z"/>`),
		g.Group(children),
	)
}