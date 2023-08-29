package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M166 346V32h-44v90H32v44h90v224h224v90h44v-90h90v-44H166z"/><path fill="currentColor" d="M346 320h44V122H192v44h154v154z"/>`),
		g.Group(children),
	)
}