package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aviation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M32.5 30H36L44 22H11C10.3509 22 9.71929 22.2105 9.2 22.6L4 26.5L10.2767 29.6767C10.6967 29.8893 11.1607 30 11.6314 30H13"/><path fill="#2F88FF" d="M27 31C27 32.6569 23.4183 34 19 34L19 28C23.4183 28 27 29.3431 27 31Z"/><path fill="#2F88FF" stroke-linecap="round" d="M31 18L22 22H37L40 12H37L31 18Z"/></g>`),
		g.Group(children),
	)
}