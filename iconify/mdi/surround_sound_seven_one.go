package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSoundSevenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17h-2v-2h2v2m6-10v10h-2V9h-2V7h4M4 17l4-8H4V7h6v2l-4 8"/>`),
		g.Group(children),
	)
}