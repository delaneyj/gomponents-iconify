package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M14 21H5V5H21V14"/><path stroke-linecap="round" d="M32 27H43V43H27V32"/><path fill="#2F88FF" d="M14 32V14H32V32H14Z"/></g>`),
		g.Group(children),
	)
}