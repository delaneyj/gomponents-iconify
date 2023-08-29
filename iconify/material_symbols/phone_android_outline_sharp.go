package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneAndroidOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20h4v-1h-4v1Zm-5 3V1h14v22H5Zm2-5v3h10v-3H7Zm0-2h10V6H7v10ZM7 4h10V3H7v1Zm0 14v3v-3ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}