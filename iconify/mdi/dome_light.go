package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10a9 9 0 0 1-18 0v-.97c0-.27.11-.53.3-.73c.2-.19.46-.3.73-.3H20c.25 0 .5.11.7.3c.19.2.3.45.3.7v1M5 10c0 3.86 3.14 7 7 7s7-3.14 7-7H5m15-4v1H4V6c0-.27.1-.5.29-.71C4.5 5.1 4.73 5 5 5h5V3h4v2h5c.25 0 .5.11.7.3c.19.2.3.45.3.7Z"/>`),
		g.Group(children),
	)
}