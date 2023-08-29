package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M42 21H6V42H42V21Z"/><path stroke-linecap="round" d="M14 21V6H34V21"/><path stroke-linecap="round" d="M20 12V14"/><path stroke-linecap="round" d="M28 12V14"/></g>`),
		g.Group(children),
	)
}