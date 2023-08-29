package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OffScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33 6V15H42"/><path d="M15 6V15H6"/><path d="M15 42V33H6"/><path d="M33 42V33H41.8995"/></g>`),
		g.Group(children),
	)
}