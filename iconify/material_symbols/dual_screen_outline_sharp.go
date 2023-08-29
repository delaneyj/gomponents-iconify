package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 17.6l6 2.4V6.4L6 4v13.6Zm-2 1.35V2h2l8 3.025V22.95l-10-4Zm8 .05v-2h6V4H6V2h14v17h-8Zm-6-1.4V4v13.6Z"/>`),
		g.Group(children),
	)
}