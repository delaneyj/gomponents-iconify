package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3M12 3c5.5 0 10 3.58 10 8c0 .58-.08 1.14-.22 1.68A6.005 6.005 0 0 0 13 18l.08.95L12 19c-1.24 0-2.43-.18-3.53-.5C5.55 21 2 21 2 21c2.33-2.33 2.7-3.9 2.75-4.5C3.05 15.07 2 13.14 2 11c0-4.42 4.5-8 10-8Z"/>`),
		g.Group(children),
	)
}