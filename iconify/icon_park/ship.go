package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 20.3768L24 14L42 20.3768L36.6667 34H11.3333L6 20.3768Z" clip-rule="evenodd"/><path fill="#2F88FF" d="M13 8H35L34.9976 17.8961L24 14L13 17.8958V8Z"/><path stroke-linecap="round" d="M24 8V4"/><path stroke-linecap="round" d="M24 24V16"/><path stroke-linecap="round" d="M10 40L13.5 44L17 40L20.5 44L24 40L27.5 44L31 40L34.5 44L38 40"/></g>`),
		g.Group(children),
	)
}