package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReverseOperationIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="18" height="32" x="6" y="5" fill="#2F88FF" stroke="#000"/><rect width="18" height="32" x="24" y="11" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M13 17L17 20.7895L13.3333 25"/><path stroke="#fff" d="M35 23L31 26.7895L34.6667 31"/></g>`),
		g.Group(children),
	)
}