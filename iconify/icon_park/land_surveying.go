package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandSurveying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 26V44"/><path stroke-linecap="round" d="M24 26L36 44"/><path stroke-linecap="round" d="M24 26L12 44"/><path stroke-linecap="round" d="M24 14V20"/><path stroke-linecap="round" d="M19 20L29 20"/><rect width="26" height="8" x="10" y="6" fill="#2F88FF"/><path stroke-linecap="round" d="M40 8V12"/></g>`),
		g.Group(children),
	)
}