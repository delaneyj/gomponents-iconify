package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 30L9 6H39L44 30"/><path fill="#2F88FF" d="M4 30H14.9091L16.7273 36H31.2727L33.0909 30H44V43H4V30Z"/><path stroke-linecap="round" d="M18 20L24 14L30 20"/><path stroke-linecap="round" d="M24 26V14"/></g>`),
		g.Group(children),
	)
}