package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.984 9.261V.735c0-.371-.371-.672-.83-.672h-1.31c-.458 0-.829.301-.829.672V9.26c-1.273.519-2.016 1.766-2.016 3.225a3.485 3.485 0 0 0 6.968 0c.002-1.459-.711-2.705-1.983-3.224zM7.947.957H9.02V10H7.947V.957z"/>`),
		g.Group(children),
	)
}