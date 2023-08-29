package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppletClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m42 42l-9-9M6 6l9 9l-9-9Zm0 36l9-9M42 6l-9 9l9-9ZM24 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`),
		g.Group(children),
	)
}