package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 4H32.8889L40 11.2727V44H8V4Z"/><path fill="#43CCF8" stroke="#fff" d="M33 26H15V36H33V26Z"/><path stroke="#fff" stroke-linecap="round" d="M15 12V18"/><path stroke="#fff" stroke-linecap="round" d="M21 12V18"/><path stroke="#fff" stroke-linecap="round" d="M27 12V18"/></g>`),
		g.Group(children),
	)
}