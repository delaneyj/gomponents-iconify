package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M13 4H35L34 13L25.5 19H22.5L14 13L13 4Z"/><path d="M24 19V30"/><rect width="6" height="14" x="21" y="30" fill="#2F88FF" rx="3"/></g>`),
		g.Group(children),
	)
}