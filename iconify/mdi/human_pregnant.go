package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanPregnant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 4c0-1.11.89-2 2-2c1.11 0 2 .89 2 2c0 1.11-.89 2-2 2c-1.11 0-2-.89-2-2m7 9c0-1.34-.83-2.5-2-3a3 3 0 0 0-3-3a3 3 0 0 0-3 3v7h2v5h3v-5h3v-4Z"/>`),
		g.Group(children),
	)
}