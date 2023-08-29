package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneAndroidSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20h4v-1h-4v1Zm-5 3V1h14v22H5Zm2-7h10V6H7v10Z"/>`),
		g.Group(children),
	)
}