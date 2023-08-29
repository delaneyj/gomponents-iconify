package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h4v4H6V2zM3 7h10v9h-2v6H9v-6H7v6H5v-6H3V7zm18-4h-6v2h6V3zm-4 4h4v2h-4V7zm4 4h-6v2h6v-2zm-6 8h6v2h-6v-2zm6-4h-4v2h4v-2z"/>`),
		g.Group(children),
	)
}