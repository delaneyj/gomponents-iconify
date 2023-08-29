package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVipOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m4.503 16.366l8.013-8.694A2.13 2.13 0 0 1 14.082 7h19.836a2.13 2.13 0 0 1 1.566.672l8.013 8.694a1.85 1.85 0 0 1-.035 2.572L24.75 40.15a1 1 0 0 1-1.5 0L4.538 18.938a1.85 1.85 0 0 1-.035-2.572Z"/><path d="m16 20l8 9l8-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVipOne0)"/>`),
		g.Group(children),
	)
}