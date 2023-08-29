package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baokemeng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBaokemeng0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" d="M44 24H30c0-3.31-2.69-6-6-6s-6 2.69-6 6H4C4 12.95 12.95 4 24 4s20 8.95 20 20Z"/><path stroke-linecap="round" d="M18 24H4c0 11.05 8.95 20 20 20s20-8.95 20-20H30"/><path d="M24 30a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBaokemeng0)"/>`),
		g.Group(children),
	)
}