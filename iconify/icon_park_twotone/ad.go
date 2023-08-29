package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAd0"><g fill="none"><circle cx="24" cy="24" r="20" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m38 38l-3-3M10 10l3 3"/><path fill="#555" d="M21.143 28L18 17l-3.143 11h6.286Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m14 31l.857-3M22 31l-.857-3m0 0L18 17l-3.143 11m6.286 0h-6.286"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 24c0 5-3.582 7-8 7V17c4.418 0 8 2 8 7Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAd0)"/>`),
		g.Group(children),
	)
}