package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanHandsdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2h4v4h-4V2zM7 7h10v2h-2v13h-2v-6h-2v6H9V9H7V7zm-2 4h2V9H5v2zm0 0v2H3v-2h2zm14 0h-2V9h2v2zm0 0h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}