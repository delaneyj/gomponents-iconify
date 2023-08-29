package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperModeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.6 16.6L4 12l4.6-4.6L10 8.85L6.85 12L10 15.15L8.6 16.6ZM5 17h2v1h10v-1h2v6H5v-6ZM7 7H5V1h14v6h-2V6H7v1Zm0 13v1h10v-1H7ZM7 4h10V3H7v1Zm8.4 12.6L14 15.15L17.15 12L14 8.85l1.4-1.45L20 12l-4.6 4.6ZM7 4V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}