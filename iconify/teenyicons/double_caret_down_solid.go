package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.707 1.474L7.5 9.234L.293 1.475l.733-.68L7.5 7.764L13.974.793l.733.68Zm-13.68 4.82l6.473 6.97l6.474-6.972l.733.68L7.5 14.736L.293 6.974l.733-.68Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}