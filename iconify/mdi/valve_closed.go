package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValveClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v20h-2v-9h-5.18a3 3 0 0 1-5.64 0H4v9H2V2h2v9h5.18a3 3 0 0 1 5.64 0H20V2Z"/>`),
		g.Group(children),
	)
}