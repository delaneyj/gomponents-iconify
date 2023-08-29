package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cuvette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M18 10H30V27L44 41V44H4V41L18 27V10Z" clip-rule="evenodd"/><path stroke="#000" d="M40.5 37.5L37 34L33.5 30.5"/><path stroke="#000" d="M40.5 37.5L37 34L33.5 30.5"/><path stroke="#000" d="M14.5 30.5L11 34L7.5 37.5"/><path stroke="#000" d="M14.5 30.5L11 34L7.5 37.5"/><path stroke="#000" d="M18 4H30"/><path stroke="#fff" d="M24 27V28"/><path stroke="#fff" d="M24 18V21"/><path stroke="#000" d="M10 35H38"/></g>`),
		g.Group(children),
	)
}