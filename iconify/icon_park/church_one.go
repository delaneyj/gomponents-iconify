package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M24 4V16"/><path stroke="#000" stroke-linecap="round" d="M20 8L28 8"/><path stroke="#000" stroke-linejoin="round" d="M15 28H9C7.89543 28 7 28.8954 7 30V44"/><path stroke="#000" stroke-linejoin="round" d="M33 28H39C40.1046 28 41 28.8954 41 30V44"/><path stroke="#000" stroke-linecap="round" d="M4 44L44 44"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M15 23L24 15L33 23V44H15V23Z"/><path stroke="#fff" stroke-linecap="round" d="M24 34V44"/><path stroke="#000" stroke-linecap="round" d="M20 44L28 44"/></g>`),
		g.Group(children),
	)
}