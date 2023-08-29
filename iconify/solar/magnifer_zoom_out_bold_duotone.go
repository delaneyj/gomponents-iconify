package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagniferZoomOutBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.157 20.313a9.157 9.157 0 1 0 0-18.313a9.157 9.157 0 0 0 0 18.313Z" opacity=".5"/><path d="M17.1 18.122a9.206 9.206 0 0 0 1.022-1.022l3.666 3.666a.723.723 0 0 1-1.022 1.022L17.1 18.122Z"/><path fill-rule="evenodd" d="M8.023 11.156c0-.399.324-.722.723-.722h4.82a.723.723 0 1 1 0 1.445h-4.82a.723.723 0 0 1-.723-.723Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}