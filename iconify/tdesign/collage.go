package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h4.198l3.555-16H4Zm9.802 0l-1.641 7.387L20 14.523V4h-6.198ZM20 16.677l-8.279-3.312L10.247 20H20v-3.323Z"/>`),
		g.Group(children),
	)
}