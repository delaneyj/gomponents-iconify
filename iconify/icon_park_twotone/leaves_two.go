package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeavesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLeavesTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M21 17c8.385-2.12 17.665-8.76 21-12c0 15-3.801 23.472-6 26c-10 11.5-20.935 6.16-23 2c-4.855-9.777 2.07-14.5 8-16Z"/><path d="M6 43c.412-2 2.388-6.6 7-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLeavesTwo0)"/>`),
		g.Group(children),
	)
}