package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToBackSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9V7h2v2H7Zm0 4v-2h2v2H7Zm0-8V3h2v2H7Zm4 12v-2h2v2h-2Zm8-12V3h2v2h-2Zm-8 0V3h2v2h-2ZM7 17v-2h2v2H7Zm12-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0 8v-2h2v2h-2ZM3 21V7h2v12h12v2H3ZM15 5V3h2v2h-2Zm0 12v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}