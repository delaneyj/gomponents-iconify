package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M33 5H5V43H33V5Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M33 21H43V43H33"/><path stroke="#fff" stroke-linecap="round" d="M13 21H25"/><path stroke="#fff" stroke-linecap="round" d="M19 15V27"/></g>`),
		g.Group(children),
	)
}