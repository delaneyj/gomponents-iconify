package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 38V42L6 42L6 38"/><path d="M30 6L40 16L30 26"/><path d="M40 16C20 16 6 19 6 32"/></g>`),
		g.Group(children),
	)
}