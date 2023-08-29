package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M16 15C16 13.3431 17.3431 12 19 12H37C38.6569 12 40 13.3431 40 15V33C40 34.6569 38.6569 36 37 36H19C17.3431 36 16 34.6569 16 33V15Z"/><path stroke="#000" stroke-linecap="round" d="M8 4L8 44"/><path stroke="#000" stroke-linecap="round" d="M8 19L16 19"/><path stroke="#000" stroke-linecap="round" d="M8 29L16 29"/><path stroke="#fff" stroke-linecap="round" d="M22 24L34 24"/><path stroke="#fff" stroke-linecap="round" d="M28 18V30"/></g>`),
		g.Group(children),
	)
}