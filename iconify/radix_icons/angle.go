package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.891 2.194a.5.5 0 0 1 .115.697L2.474 12H13.5a.5.5 0 0 1 0 1h-12a.5.5 0 0 1-.406-.791l7.1-9.9a.5.5 0 0 1 .697-.115ZM11.1 6.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM10.4 4a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.7 4.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm1.3 1.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}