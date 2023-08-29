package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M344 16H168v152H16v176h152v152h176V344h152V168H344Zm120 184v112H312v152H200V312H48V200h152V48h112v152Z"/>`),
		g.Group(children),
	)
}