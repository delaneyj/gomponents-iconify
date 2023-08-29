package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyScreenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9.6L7.6 6H4Zm0 7L14.575 6H10.4L4 12.425ZM5.4 18H20V6h-2.6ZM2 20V4h20v16Z"/>`),
		g.Group(children),
	)
}