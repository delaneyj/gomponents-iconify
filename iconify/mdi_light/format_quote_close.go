package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatQuoteClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6v8l-1.95 4h-4.2l1.95-4H12V6h6m-1 7.77V7h-4v6h2.4l-1.95 4h1.97L17 13.77M11 6v8l-1.95 4h-4.2l1.95-4H5V6h6m-1 7.77V7H6v6h2.4l-1.95 4h1.97L10 13.77Z"/>`),
		g.Group(children),
	)
}