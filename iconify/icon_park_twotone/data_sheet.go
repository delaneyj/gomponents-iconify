package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDataSheet0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" fill="#555" rx="2"/><path d="M32 25v7m-8-16v16m-8-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDataSheet0)"/>`),
		g.Group(children),
	)
}