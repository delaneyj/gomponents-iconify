package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValveOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22H2V2h2m18 0h-2v20h2M11 4v5.18a3 3 0 0 0 0 5.64V20h2v-5.18a3 3 0 0 0 0-5.64V4Z"/>`),
		g.Group(children),
	)
}