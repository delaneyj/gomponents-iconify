package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanCircleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48C141.31 48 48 141.31 48 256s93.31 208 208 208s208-93.31 208-208S370.69 48 256 48Zm-24 320h-44a44.05 44.05 0 0 1-44-44v-44h32v44a12 12 0 0 0 12 12h44Zm0-192h-44a12 12 0 0 0-12 12v44h-32v-44a44.05 44.05 0 0 1 44-44h44Zm136 148a44.05 44.05 0 0 1-44 44h-44v-32h44a12 12 0 0 0 12-12v-44h32Zm0-92h-32v-44a12 12 0 0 0-12-12h-44v-32h44a44.05 44.05 0 0 1 44 44Z"/>`),
		g.Group(children),
	)
}