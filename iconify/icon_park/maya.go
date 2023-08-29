package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="40" height="8" x="4" y="34" stroke-linejoin="round"/><rect width="34" height="7" x="7" y="27" stroke-linejoin="round"/><rect width="28" height="6" x="10" y="21" stroke-linejoin="round"/><rect width="22" height="6" x="13" y="15" stroke-linejoin="round"/><rect width="8" height="8" x="20" y="7" stroke-linejoin="round"/><path d="M20 15L14 42"/><path d="M28 15L34 42"/></g>`),
		g.Group(children),
	)
}