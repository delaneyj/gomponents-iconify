package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChickenZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 31C30 26.5 31.4 17.6 41 14"/><path d="M11 4L15 10"/><path d="M38 31C38 33.8889 34.24 38.5556 28 40L26 44"/><path d="M21 10.2738C15.1845 9.2518 3.86133 10.5366 5.09285 23.8519C5.09286 28.2319 7.37935 37.1968 16 40L18 44"/><path d="M19 18C19.6957 21.8333 22.5 31 31 31H38C37.3043 28.6667 37.1565 23.4 43 21"/></g>`),
		g.Group(children),
	)
}