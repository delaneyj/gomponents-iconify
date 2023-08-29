package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcasePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3M10 2h4a2 2 0 0 1 2 2v2h4a2 2 0 0 1 2 2v5.53A5.97 5.97 0 0 0 18 12a6 6 0 0 0-6 6c0 1.09.29 2.12.8 3H4a2 2 0 0 1-2-2V8c0-1.11.89-2 2-2h4V4c0-1.11.89-2 2-2m4 4V4h-4v2h4Z"/>`),
		g.Group(children),
	)
}