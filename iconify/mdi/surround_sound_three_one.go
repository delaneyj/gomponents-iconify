package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSoundThreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17h-2v-2h2v2m6-10v10h-2V9h-2V7h4m-10 8c0 1.1-.9 2-2 2H4v-2h4v-2H6v-2h2V9H4V7h4c1.1 0 2 .9 2 2v1.5c0 .8-.7 1.5-1.5 1.5c.8 0 1.5.7 1.5 1.5V15"/>`),
		g.Group(children),
	)
}