package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 4H12V44H42V4Z"/><path stroke="#000" stroke-linecap="round" d="M4 11.7895L12 10V44L4 11.7895Z" clip-rule="evenodd"/><path fill="#43CCF8" stroke="#fff" d="M27 18L22 24L27 30L32 24L27 18Z"/><path stroke="#fff" stroke-linecap="round" d="M18 10V14"/><path stroke="#fff" stroke-linecap="round" d="M36 34V38"/></g>`),
		g.Group(children),
	)
}