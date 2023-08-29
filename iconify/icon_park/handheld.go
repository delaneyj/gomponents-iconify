package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handheld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 18V44H6V18"/><path fill="#2F88FF" d="M42 4H6V18H42V4Z"/><path stroke-linecap="round" d="M16 27V35"/><path stroke-linecap="round" d="M12 31H20"/><path fill="#2F88FF" d="M32 35C34.2091 35 36 33.2091 36 31C36 28.7909 34.2091 27 32 27C29.7909 27 28 28.7909 28 31C28 33.2091 29.7909 35 32 35Z"/></g>`),
		g.Group(children),
	)
}