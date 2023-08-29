package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.12 8.46l5.66 5.66l-5.66 5.66l-5.66-5.66l-4.24 4.24V4.22h14.14l-4.24 4.24m-7.78 4.95l2.12-2.12l5.66 5.66l2.83-2.83l-5.66-5.65l2.12-2.13H6.34v7.07Z"/>`),
		g.Group(children),
	)
}