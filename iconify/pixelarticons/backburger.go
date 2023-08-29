package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7h10v2H11V7zm-8 4h2V9h2v2h14v2H7v2H5v-2H3v-2zm4 4v2h2v-2H7zm0-6V7h2v2H7zm14 6H11v2h10v-2z"/>`),
		g.Group(children),
	)
}