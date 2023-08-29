package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9h-2v3H8v2h3v3h2v-3h3v-2h-3m-1 8a7 7 0 0 1-7-7a7 7 0 0 1 7-7a7 7 0 0 1 7 7a7 7 0 0 1-7 7m0-16a9 9 0 0 0-9 9a9 9 0 0 0 9 9a9 9 0 0 0 9-9a9 9 0 0 0-9-9m10 1.72l-4.6-3.86l-1.29 1.53l4.6 3.86M7.88 3.39L6.6 1.86L2 5.71l1.29 1.53l4.59-3.85Z"/>`),
		g.Group(children),
	)
}