package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intercom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M13 14C13 12.3431 14.3431 11 16 11H32C33.6569 11 35 12.3431 35 14V25L33 31V41C33 42.6569 31.6569 44 30 44H18C16.3431 44 15 42.6569 15 41V31L13 25V14Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M19 11L19 4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 11L28 7"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M28 19L20 19"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M26 26L22 26"/></g>`),
		g.Group(children),
	)
}