package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21.725v-9.15L3 7.95v9.175l8 4.6Zm2 0l8-4.6V7.95l-8 4.625v9.15Zm3.975-13.75l2.95-1.725L12 1.7L9.025 3.4l7.95 4.575ZM12 10.85l2.975-1.7l-7.925-4.6l-3 1.725L12 10.85Z"/>`),
		g.Group(children),
	)
}