package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="30" height="38" x="9" y="5" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 18h30"/><circle cx="24" cy="30" r="6"/></g>`),
		g.Group(children),
	)
}