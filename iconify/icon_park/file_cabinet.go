package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCabinet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M42 4H6V14H42V4Z"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M42 19H6V29H42V19Z"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M42 34H6V44H42V34Z"/><path stroke="#fff" stroke-linecap="round" d="M21 9H27"/><path stroke="#fff" stroke-linecap="round" d="M21 24H27"/><path stroke="#fff" stroke-linecap="round" d="M21 39H27"/></g>`),
		g.Group(children),
	)
}