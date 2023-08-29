package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BezierCurve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="8" height="8" x="4" y="30" fill="#2F88FF"/><rect width="8" height="8" x="36" y="30" fill="#2F88FF"/><rect width="8" height="8" x="20" y="10" fill="#2F88FF"/><path stroke-linecap="round" d="M20 14H4"/><path stroke-linecap="round" d="M44 14H28"/><path stroke-linecap="round" d="M8 30.0001C8 22.5447 13.0991 16.2803 20 14.5042"/><path stroke-linecap="round" d="M28 14.5042C34.9009 16.2803 40 22.5447 40 30.0001"/></g>`),
		g.Group(children),
	)
}