package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M8 20V38C8 39.1046 8.89543 40 10 40H38C39.1046 40 40 39.1046 40 38V20"/><path fill="#2F88FF" d="M5 14H43V20H5V14Z"/><path fill="#2F88FF" d="M31 8H17V14H31V8Z"/><path stroke-linecap="round" d="M20 30L28 30"/><path stroke-linecap="round" d="M24 26V34"/></g>`),
		g.Group(children),
	)
}