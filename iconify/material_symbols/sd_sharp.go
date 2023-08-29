package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm4-5h5v-3.5H7.5v-1h2v.5H11V9H6v3.5h3.5v1h-2V13H6v2Zm7 0h4.25l.75-.75v-4.5L17.25 9H13v6Zm1.5-1.5v-3h2v3h-2Z"/>`),
		g.Group(children),
	)
}