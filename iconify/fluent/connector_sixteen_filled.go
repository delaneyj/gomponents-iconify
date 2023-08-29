package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectorSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 3a.5.5 0 0 0-.5.5v1.882l-.947 1.894A.5.5 0 0 0 2 7.5v3a.5.5 0 0 0 .5.5H3v1.5a.5.5 0 0 0 1 0V11h1v1.5a.5.5 0 0 0 1 0V11h.5a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.053-.224L6 5.382V3.5a.5.5 0 0 0-.5-.5h-2Zm9 10h-2a.5.5 0 0 1-.5-.5v-1.882l-.947-1.894A.5.5 0 0 1 9 8.5v-3a.5.5 0 0 1 .5-.5h.5V3.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5V5h.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.053.224L13 10.618V12.5a.5.5 0 0 1-.5.5ZM11 5h1V4h-1v1Z"/>`),
		g.Group(children),
	)
}