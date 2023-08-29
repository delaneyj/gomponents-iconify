package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VitalSignsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 13v-2h5.7L9 17.1l6-15.825L18.7 11H23v2h-5.7L15 6.9L9 22.725L5.3 13H1Z"/>`),
		g.Group(children),
	)
}