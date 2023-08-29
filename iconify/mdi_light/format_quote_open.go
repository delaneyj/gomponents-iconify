package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatQuoteOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18v-8l1.95-4h4.2L9.2 10H11v8H5m1-7.77V17h4v-6H7.6l1.95-4H7.58L6 10.23M12 18v-8l1.95-4h4.2l-1.95 4H18v8h-6m1-7.77V17h4v-6h-2.4l1.95-4h-1.97L13 10.23Z"/>`),
		g.Group(children),
	)
}