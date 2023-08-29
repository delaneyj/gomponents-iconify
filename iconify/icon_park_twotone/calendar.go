package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCalendar0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5 19h38v21a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V19Z"/><path d="M5 9a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v10H5V9Z"/><path stroke-linecap="round" d="M16 4v8m16-8v8m-4 22h6m-20 0h6m8-8h6m-20 0h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCalendar0)"/>`),
		g.Group(children),
	)
}