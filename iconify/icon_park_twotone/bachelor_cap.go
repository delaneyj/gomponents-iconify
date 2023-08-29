package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BachelorCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBachelorCap0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m4 13l20-5l20 5l-20 5l-20-5Z"/><path d="M13 16v9.97S18 29 24 29s11-3.03 11-3.03V16M7 14v22"/><path fill="#555" d="M4 34h6v6H4z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBachelorCap0)"/>`),
		g.Group(children),
	)
}