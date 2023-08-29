package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M6 6H42V42H6"/><path fill="#2F88FF" stroke="#000" d="M6 6V42L24 36V12L6 6Z"/><path stroke="#fff" d="M18 22V26"/></g>`),
		g.Group(children),
	)
}