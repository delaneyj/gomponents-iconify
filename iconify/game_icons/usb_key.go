package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m394.8 30.88l-65 65.03l86.3 86.29l65.1-65l-86.4-86.32zm-6.3 36.04l17 17l-12.8 12.72l-17-17l12.8-12.72zm-82.8 30.4l-11.3 11.28l109 108.9l11.3-11.2l-109-108.98zM263.3 103L23.4 342.9v60.5l85.2 85.2h60.5l240-239.9L263.3 103zm164.9 3.6l16.9 17l-12.8 12.6l-16.9-17l12.8-12.6z"/>`),
		g.Group(children),
	)
}