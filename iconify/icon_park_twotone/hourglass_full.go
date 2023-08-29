package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHourglassFull0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M7 4h34M7 44h34"/><path fill="#555" d="M11 44c2.667-13.339 7-20.006 13-20c6 .006 10.333 6.672 13 20H11Z"/><path fill="#555" d="M37 4c-2.667 13.339-7 20.006-13 20c-6-.006-10.333-6.672-13-20h26Z"/><path stroke-linecap="round" d="M21 15h6m-8 23h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHourglassFull0)"/>`),
		g.Group(children),
	)
}