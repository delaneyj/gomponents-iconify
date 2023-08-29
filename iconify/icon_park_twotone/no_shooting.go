package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoShooting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNoShooting0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linejoin="round" d="m15 12l3-6h12l3 6H15Z"/><path fill="#555" stroke-linejoin="round" d="M41 12H7c-1.657 0-3 1.254-3 2.8v24.4C4 40.746 5.343 42 7 42h34c1.657 0 3-1.254 3-2.8V14.8c0-1.546-1.343-2.8-3-2.8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m28 20l-8 14"/><circle cx="24" cy="27" r="8" stroke-linecap="round"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNoShooting0)"/>`),
		g.Group(children),
	)
}