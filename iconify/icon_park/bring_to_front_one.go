package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringToFrontOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 21V34H27"/><path d="M21 14H34V27"/><path fill="#2F88FF" d="M5 21V5H21V21H5Z"/><path fill="#2F88FF" d="M27 43V27H43V43H27Z"/></g>`),
		g.Group(children),
	)
}