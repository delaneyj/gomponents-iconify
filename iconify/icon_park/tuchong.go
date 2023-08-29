package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuchong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M5 39H43V28H37V33H11V15H23V9H5V39Z"/><path stroke-linecap="round" d="M43 16V22C35 22 29 17 29 9H35C35 13 37 16 43 16Z"/></g>`),
		g.Group(children),
	)
}