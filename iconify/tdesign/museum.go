package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Museum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h12v4.327l4-.444V22H2V8.105l4-.444V2Zm2 5.438l8-.889V4H8v3.438ZM18 20h2v-5h-2v5Zm2-7V8.117L4 9.895V20h12v-7h4ZM6 10.998h2.004v2.004H6v-2.004Zm4 0h2.004v2.004H10v-2.004Z"/>`),
		g.Group(children),
	)
}