package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RsMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRsMale0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4.31 16.27A2 2 0 0 1 6.29 14h35.42a2 2 0 0 1 1.981 2.27l-2.454 18A2 2 0 0 1 39.254 36H8.746a2 2 0 0 1-1.982-1.73l-2.454-18Z"/><path d="M19 22h2m2 6h2m-14-6h2m2 6h2m10-6h2m2 6h2m2-6h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRsMale0)"/>`),
		g.Group(children),
	)
}