package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h10v-2H9v2zm0-6h6v-2H9v2zm0-8v2h12V5H9zm-4-.5a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 5 4.5zm0 6a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 5 10.5zm0 6a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 5 16.5z"/>`),
		g.Group(children),
	)
}