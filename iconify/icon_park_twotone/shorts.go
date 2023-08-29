package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shorts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShorts0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 6H11v10L4 33l15 9l5-6l5 6l15-9l-7-17V6Z"/><path fill="#fff" d="M11 14a2 2 0 1 0 0 4v-4Zm26 4a2 2 0 1 0 0-4v4Zm-26 0h26v-4H11v4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 14v2l1.75 4.25M11 14v2l-1.75 4.25"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShorts0)"/>`),
		g.Group(children),
	)
}