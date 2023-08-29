package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L2 12h3v8h14v-8h3L12 3m0 5.75A2.25 2.25 0 0 1 14.25 11A2.25 2.25 0 0 1 12 13.25A2.25 2.25 0 0 1 9.75 11A2.25 2.25 0 0 1 12 8.75M12 15c1.5 0 4.5.75 4.5 2.25V18h-9v-.75c0-1.5 3-2.25 4.5-2.25Z"/>`),
		g.Group(children),
	)
}