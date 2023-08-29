package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLayers0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M4 11.914L24 19l20-7.086L24 5L4 11.914Z"/><path stroke-linecap="round" d="m4 20l20 7l20-7M4 28l20 7l20-7M4 36l20 7l20-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLayers0)"/>`),
		g.Group(children),
	)
}