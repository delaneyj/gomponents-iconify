package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6h5v6h-5V6m-6 6V6h5v6H9m6 7v-6h5v6h-5m-6 0v-6h5v6H9m-6 0v-6h5v6H3m0-7V6h5v6H3m1-5v4h3V7H4m6 0v4h3V7h-3m6 0v4h3V7h-3M4 14v4h3v-4H4m6 0v4h3v-4h-3m6 0v4h3v-4h-3Z"/>`),
		g.Group(children),
	)
}