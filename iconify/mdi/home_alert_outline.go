package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 5.7l5 4.5V18H7v-7.8l5-4.5M19 20v-8h3L12 3L2 12h3v8m8-12h-2v5h2V8m0 7h-2v2h2v-2"/>`),
		g.Group(children),
	)
}