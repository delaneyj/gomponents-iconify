package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterKeyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M23 23V5H43V43H5V23H23Z"/><path stroke="#fff" d="M33 13V33H13"/><path stroke="#fff" d="M17 29L13 33L17 37"/></g>`),
		g.Group(children),
	)
}