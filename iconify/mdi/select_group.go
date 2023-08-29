package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3a2 2 0 0 0-2 2h2m2-2v2h2V3m2 0v2h2V3m2 0v2h2V3m2 0v2h2a2 2 0 0 0-2-2M3 7v2h2V7m2 0v4h4V7m2 0v4h4V7m2 0v2h2V7M3 11v2h2v-2m14 0v2h2v-2M7 13v4h4v-4m2 0v4h4v-4M3 15v2h2v-2m14 0v2h2v-2M3 19a2 2 0 0 0 2 2v-2m2 0v2h2v-2m2 0v2h2v-2m2 0v2h2v-2m2 0v2a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}