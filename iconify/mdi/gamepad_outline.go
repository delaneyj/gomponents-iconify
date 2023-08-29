package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 9H2v6h5.5l3-3l-3-3M6 13H4v-2h2v2m9-5.5V2H9v5.5l3 3l3-3M11 4h2v2h-2V4M9 16.5V22h6v-5.5l-3-3l-3 3m4 3.5h-2v-2h2v2m3.5-11l-3 3l3 3H22V9h-5.5m3.5 4h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}