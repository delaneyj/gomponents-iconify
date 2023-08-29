package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weightlifting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWeightlifting0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#555" d="M24 27a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 9h40M4 4v10M44 4v10M11 9v17.1L24 34l13-8V9M24 34v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWeightlifting0)"/>`),
		g.Group(children),
	)
}