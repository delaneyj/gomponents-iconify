package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOctagon0"><path fill="#555" stroke="#fff" stroke-width="4" d="M15.41 42.389L5.56 32.18A2 2 0 0 1 5 30.792V17.208a2 2 0 0 1 .56-1.39l9.85-10.207A2 2 0 0 1 16.85 5h14.3a2 2 0 0 1 1.44.611l9.85 10.208a2 2 0 0 1 .56 1.389v13.584a2 2 0 0 1-.56 1.39l-9.85 10.207a2 2 0 0 1-1.44.611h-14.3a2 2 0 0 1-1.44-.611Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOctagon0)"/>`),
		g.Group(children),
	)
}