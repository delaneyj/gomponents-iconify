package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInspection0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M43 33V19H5v22a2 2 0 0 0 2 2h17"/><path stroke-linejoin="round" d="M5 10a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v9H5v-9Z"/><path stroke-linecap="round" d="M16 5v8m16-8v8"/><circle cx="30" cy="32" r="7" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m36 37l6 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInspection0)"/>`),
		g.Group(children),
	)
}