package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Makeups(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMakeups0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4"/><path d="m37.61 9.472l.255.786h.827l-.669.486l.255.786l-.668-.486l-.669.486l.255-.786l-.668-.486h.826l.256-.786Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 31s2 4 8 4s8-4 8-4"/><circle cx="17" cy="22" r="3" fill="#fff"/><circle cx="31" cy="22" r="3" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMakeups0)"/>`),
		g.Group(children),
	)
}