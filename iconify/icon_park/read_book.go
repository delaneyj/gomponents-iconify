package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M24 21V44C20.7086 40 10.6286 40 6 40V18C15.8743 18 22.1143 20 24 21Z"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M24 21V44C27.2914 40 37.3714 40 42 40V18C32.1257 18 25.8857 20 24 21Z"/><circle cx="24" cy="12" r="8"/></g>`),
		g.Group(children),
	)
}