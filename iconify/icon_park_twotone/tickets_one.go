package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTicketsOne0"><g fill="none"><rect width="26" height="38" x="5" y="42" fill="#555" stroke="#fff" stroke-linejoin="bevel" stroke-width="4" rx="2" transform="rotate(-90 5 42)"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 16L32 5l5 11"/><circle cx="13" cy="23" r="2" fill="#fff"/><circle cx="13" cy="29" r="2" fill="#fff"/><circle cx="13" cy="35" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 35h4l11-12m-12 6h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTicketsOne0)"/>`),
		g.Group(children),
	)
}