package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChair0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 23l-8 21m16-21l8 21M16 34h16"/><path fill="#555" d="M29.454 23H18.545C15.819 23 14 21.333 14 18.833V12h4V6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v6h4v6.833c0 2.5-1.818 4.167-4.546 4.167Z"/><path d="M30 12h8m-20 0h-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChair0)"/>`),
		g.Group(children),
	)
}