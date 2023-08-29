package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUDisk0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M30 12V4H18v8"/><path fill="#555" d="M13 12.373c0-.206.167-.373.373-.373h21.254c.206 0 .373.167.373.373V33c0 6.075-4.925 11-11 11s-11-4.925-11-11V12.373Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 21h8m-8 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUDisk0)"/>`),
		g.Group(children),
	)
}