package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileFriendlyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23V1h14v6h-2V6H6v12h10v-1h2v6H4Zm2-3v1h10v-1H6ZM6 4h10V3H6v1Zm0 0V3v1Zm0 16v1v-1Zm8.95-4l-4.25-4.25l1.4-1.4l2.85 2.85l5.65-5.65l1.4 1.4L14.95 16Z"/>`),
		g.Group(children),
	)
}