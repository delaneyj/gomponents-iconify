package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSoundFiveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17h-2v-2h2v2m6-10v10h-2V9h-2V7h4M10 7v2H6v2h2c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2H4v-2h4v-2H4V7h6Z"/>`),
		g.Group(children),
	)
}