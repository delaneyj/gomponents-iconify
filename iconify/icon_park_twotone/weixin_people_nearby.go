package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinPeopleNearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWeixinPeopleNearby0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="15" cy="10" r="4"/><circle cx="33" cy="10" r="4"/><path d="M10 20h10l-2 22h-6l-2-22Zm18 0h10l-2 22h-6l-2-22Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWeixinPeopleNearby0)"/>`),
		g.Group(children),
	)
}