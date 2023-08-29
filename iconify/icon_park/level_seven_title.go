package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelSevenTitle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 8V40"/><path d="M24 8V40"/><path d="M7 24H23"/><path d="M31.998 22H41.9997L33.9936 40"/></g>`),
		g.Group(children),
	)
}