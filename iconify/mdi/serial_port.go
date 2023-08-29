package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SerialPort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3h10v2h2v3h-3v6H8V8H5V5h2V3m10 6h2v5h-2V9m-6 6h2v7h-2v-7M5 9h2v5H5V9Z"/>`),
		g.Group(children),
	)
}