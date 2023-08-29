package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 8.25564L24.0086 3L42 8.25564V19.0337C42 30.3622 34.7502 40.4194 24.0026 44.0005C13.2521 40.4195 6 30.36 6 19.0287V8.25564Z"/><path stroke="#fff" stroke-linecap="round" d="M17 19L31 19"/><path stroke="#fff" stroke-linecap="round" d="M17 25L31 25"/><path stroke="#fff" stroke-linecap="round" d="M31 19L26 14"/><path stroke="#fff" stroke-linecap="round" d="M22 30L17 25"/></g>`),
		g.Group(children),
	)
}