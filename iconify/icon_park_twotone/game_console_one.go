package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsoleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGameConsoleOne0"><g fill="none"><rect width="28" height="40" x="10" y="4" stroke="#fff" stroke-width="4" rx="2"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 12h16v10H16z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 32h8m-4 4v-8"/><rect width="4" height="4" x="27" y="33" fill="#fff" rx="2"/><rect width="4" height="4" x="30" y="26" fill="#fff" rx="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGameConsoleOne0)"/>`),
		g.Group(children),
	)
}