package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h4v4H3V3ZM1 9V1h8v8H1Zm2 4h4v4H3v-4Zm-2 6v-8h8v8H1ZM17 3h-4v4h4V3Zm-6-2v8h8V1h-8Zm2 10h-2v4h2v2h-2v2h2v-2h2v2h4v-4h-2v-2h2v-2h-4v2h-2v-2Zm4 4h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}