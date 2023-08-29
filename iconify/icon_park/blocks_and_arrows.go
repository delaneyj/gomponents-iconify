package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlocksAndArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M20 6H6V20H20V6Z"/><path fill="#2F88FF" d="M20 28H6V42H20V28Z"/><path fill="#2F88FF" d="M42 6H28V20H42V6Z"/><path d="M28 28L42 42M28 28H42H28ZM28 28V42V28Z"/></g>`),
		g.Group(children),
	)
}