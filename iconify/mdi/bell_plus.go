package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 14v3h-3v2h3v3h2v-3h3v-2h-3v-3M12 2a2 2 0 0 0-2 2a2 2 0 0 0 0 .29C7.12 5.14 5 7.82 5 11v6l-2 2v1h9.35a6 6 0 0 1-.35-2a6 6 0 0 1 6-6a6 6 0 0 1 1 .09V11c0-3.18-2.12-5.86-5-6.71A2 2 0 0 0 14 4a2 2 0 0 0-2-2m-2 19a2 2 0 0 0 2 2a2 2 0 0 0 1.65-.87a6 6 0 0 1-.84-1.13Z"/>`),
		g.Group(children),
	)
}