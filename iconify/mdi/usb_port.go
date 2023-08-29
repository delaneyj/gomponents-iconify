package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbPort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2c-1.1 0-2 .9-2 2v8H5v4l4 4v2h6v-2l4-4v-4h-1V4a2 2 0 0 0-2-2M8 4h8v8H8m1-5v2h2V7m2 0v2h2V7Z"/>`),
		g.Group(children),
	)
}