package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneRotateLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1H3a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2m0 14H3V3h6v12m12-2h-8v2h8v6H9v-1H6v1a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2m2-3l-4-2l1.91-.91A7.516 7.516 0 0 0 14 2.5V1a9 9 0 0 1 9 9Z"/>`),
		g.Group(children),
	)
}