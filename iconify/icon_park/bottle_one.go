package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M15 30C15 28.0527 15.6316 26.1579 16.8 24.6L20.4 19.8C20.7895 19.2807 21 18.6491 21 18V4H27V18C27 18.6491 27.2105 19.2807 27.6 19.8L31.2 24.6C32.3684 26.1579 33 28.0527 33 30V42C33 43.1046 32.1046 44 31 44H17C15.8954 44 15 43.1046 15 42V30Z"/><path stroke="#fff" d="M21 10L27 10"/><path stroke="#000" d="M21 12V8"/><path stroke="#000" d="M27 12V8"/></g>`),
		g.Group(children),
	)
}