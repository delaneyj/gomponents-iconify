package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 9.25564L24.0086 4L42 9.25564V20.0337C42 31.3622 34.7502 41.4194 24.0026 45.0005C13.2521 41.4195 6 31.36 6 20.0287V9.25564Z"/><path stroke="#fff" stroke-linecap="round" d="M29.5 18.4081L18.1863 29.7218"/><path stroke="#fff" stroke-linecap="round" d="M18.1863 18.4081L29.5 29.7218"/></g>`),
		g.Group(children),
	)
}