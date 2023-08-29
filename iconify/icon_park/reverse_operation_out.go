package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReverseOperationOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="18" height="32" x="6" y="5" fill="#2F88FF" stroke="#000"/><rect width="18" height="32" x="24" y="11" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M17 17L13 20.7895L16.6667 25"/><path stroke="#fff" d="M31 23L35 26.7895L31.3333 31"/></g>`),
		g.Group(children),
	)
}