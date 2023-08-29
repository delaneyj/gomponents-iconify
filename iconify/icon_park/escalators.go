package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalators(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M35 13L13 35H4V44H13L35 22H44V13H35Z"/><path d="M19 13L28 4"/><path d="M22 4H28V10"/><path d="M14 18L5 27"/><path d="M11 27H5V21"/></g>`),
		g.Group(children),
	)
}