package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolidStateDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M44 29H4V42H44V29Z"/><path fill="#fff" d="M35.5 38C36.8807 38 38 36.8807 38 35.5C38 34.1193 36.8807 33 35.5 33C34.1193 33 33 34.1193 33 35.5C33 36.8807 34.1193 38 35.5 38Z"/><path stroke="#000" stroke-linejoin="round" stroke-width="4" d="M4 28.9998L9.03837 4.99902H39.0205L44 28.9998"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M10 35.5H27"/></g>`),
		g.Group(children),
	)
}