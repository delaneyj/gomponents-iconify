package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9542 35.0457 4 24 4C12.9543 4 4 12.9542 4 24C4 35.0457 12.9543 44 24 44Z"/><path fill="#2F88FF" d="M23.9999 44C32.9552 44 40.5358 38.1142 43.0843 30H4.91553C7.46405 38.1142 15.0446 44 23.9999 44Z"/></g>`),
		g.Group(children),
	)
}