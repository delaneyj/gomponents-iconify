package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.073 0h-5.04v1.042H7.957V0H3.083l-2 16h14.922L14.073 0zM9 15H8v-3h1v3zm0-4.958H8V7h1v3.042zM7.958 4.959v-2h1v2h-1z"/>`),
		g.Group(children),
	)
}