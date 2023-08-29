package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flatten(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2h2v8h2v2h-2v2h-2v-2H9v-2h2V2zm-2 8H7V8h2v2zm6 0V8h2v2h-2zm5 6H4v2h16v-2zm-4 4H8v2h8v-2z"/>`),
		g.Group(children),
	)
}