package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinEndSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v8h-2V6H4v12h10v2H2Zm11.95-5.625L11 11.425v2.225H9V8h5.65v2H12.4l2.95 2.95l-1.4 1.425ZM19 20q-1.25 0-2.125-.875T16 17q0-1.25.875-2.125T19 14q1.25 0 2.125.875T22 17q0 1.25-.875 2.125T19 20Z"/>`),
		g.Group(children),
	)
}