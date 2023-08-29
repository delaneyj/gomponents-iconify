package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 12.5l9.4 5.88l1 .62v-5.62l8 5l1 .62V6l-9 5.62V6L1 12.5m1.89 0l7.51-4.7v9.4l-7.51-4.7m9 0l7.51-4.7v9.4l-7.51-4.7Z"/>`),
		g.Group(children),
	)
}