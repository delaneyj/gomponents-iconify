package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M5 30L10 6H38L43 30"/><path fill="#2F88FF" d="M5 30H14.9091L16.7273 36H31.2727L33.0909 30H43V43H5V30Z"/><path stroke-linecap="round" d="M18 20L24 26L30 20"/><path stroke-linecap="round" d="M24 26V14"/></g>`),
		g.Group(children),
	)
}