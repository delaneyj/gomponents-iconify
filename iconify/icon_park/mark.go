package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M11 6V42"/><path fill="#2F88FF" d="M11 9H25L32 12H39C40.1046 12 41 12.8954 41 14V31C41 32.1046 40.1046 33 39 33H32L25 30H11V9Z"/><path stroke-linecap="round" d="M7 42H15"/></g>`),
		g.Group(children),
	)
}