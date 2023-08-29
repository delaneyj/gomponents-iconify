package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="28" x="5" y="14" stroke="#000" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 14L38 6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M23 14L10 6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 20L35 26"/><rect width="4" height="4" x="33" y="32" fill="#000" rx="2"/></g>`),
		g.Group(children),
	)
}