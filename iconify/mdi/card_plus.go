package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.09 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16c1.11 0 2 .89 2 2v7.81c-.88-.51-1.91-.81-3-.81c-3.31 0-6 2.69-6 6c0 .34.04.67.09 1M18 15v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Z"/>`),
		g.Group(children),
	)
}