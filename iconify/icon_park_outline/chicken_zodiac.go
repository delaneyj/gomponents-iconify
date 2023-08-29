package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChickenZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 31c-1-4.5.4-13.4 10-17M11 4l4 6m23 21c0 2.889-3.76 7.556-10 9l-2 4m-5-33.726c-5.816-1.022-17.139.263-15.907 13.578C5.093 28.232 7.379 37.197 16 40l2 4"/><path d="M19 18c.696 3.833 3.5 13 12 13h7c-.696-2.333-.843-7.6 5-10"/></g>`),
		g.Group(children),
	)
}