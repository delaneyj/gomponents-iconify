package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMassageTable0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M14 17a2 2 0 0 1 2-2h26a2 2 0 0 1 2 2v6H14v-6Z"/><path stroke-linecap="round" d="M26 23L14 37m18-14l12 14M14 23H6m33 8H19M6 13v10m8 0v17m30-17v17M9 14l-6-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMassageTable0)"/>`),
		g.Group(children),
	)
}