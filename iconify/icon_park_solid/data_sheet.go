package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDataSheet0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M32 25v7m-8-16v16m-8-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDataSheet0)"/>`),
		g.Group(children),
	)
}