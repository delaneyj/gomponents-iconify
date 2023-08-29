package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M44 5H3.99998V17H44V5Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M3.99998 41.0301L16.1756 28.7293L22.7549 35.0301L30.7982 27L35.2786 31.368"/><path stroke="#000" stroke-linecap="round" d="M44 16.1719V42.1719"/><path stroke="#000" stroke-linecap="round" d="M3.99998 16.1719V30.1719"/><path stroke="#000" stroke-linecap="round" d="M13.0155 43H44"/><path stroke="#fff" stroke-linecap="round" d="M17 11H38"/><path stroke="#fff" stroke-linecap="round" d="M9.99998 10.9966H11"/></g>`),
		g.Group(children),
	)
}