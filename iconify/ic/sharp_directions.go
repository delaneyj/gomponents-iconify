package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDirections(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.41 12L12 1.59L1.59 11.99L12 22.41L22.41 12zM14 14.5V12h-4v3H8v-5h6V7.5l3.5 3.5l-3.5 3.5z"/>`),
		g.Group(children),
	)
}