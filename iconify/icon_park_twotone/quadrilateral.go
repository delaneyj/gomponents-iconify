package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quadrilateral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTQuadrilateral0"><path fill="#555" stroke="#fff" stroke-width="4" d="M28.038 8H7a3 3 0 0 0-3 3v26a3 3 0 0 0 3 3h32.413c2.163 0 3.616-2.22 2.748-4.203l-11.375-26A3 3 0 0 0 28.038 8Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTQuadrilateral0)"/>`),
		g.Group(children),
	)
}