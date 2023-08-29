package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkSpoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 4V44"/><path d="M8 5V15C8 20 14 20 14 20C14 20 20 20 20 15V5"/><path d="M34 20V44"/><path fill="#2F88FF" d="M40 12C40 16.4183 37.3137 20 34 20C30.6863 20 28 16.4183 28 12C28 7.58172 30.6863 4 34 4C37.3137 4 40 7.58172 40 12Z"/></g>`),
		g.Group(children),
	)
}