package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAutoFocus0"><g fill="none"><circle cx="24" cy="24" r="9" fill="#555" stroke="#fff" stroke-width="4"/><circle r="3" fill="#fff" transform="matrix(-1 0 0 1 24 24)"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 14s7.5-11.5 20.5-7S42 24.5 42 24.5M39 34s-6 11-19.5 7.5S6 24 6 24M42 8v16M6 24v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAutoFocus0)"/>`),
		g.Group(children),
	)
}