package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Log(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLog0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M13 10h28v34H13z"/><path stroke="#fff" stroke-linecap="round" d="M35 10V4H8a1 1 0 0 0-1 1v33h6"/><path stroke="#000" stroke-linecap="round" d="M21 22h12m-12 8h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLog0)"/>`),
		g.Group(children),
	)
}