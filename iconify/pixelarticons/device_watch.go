package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceWatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2h8v4h5v12h-5v4H8v-4H3V6h5V2zM5 16h14V8H5v8zm6-6h2v2h2v2h-4v-4z"/>`),
		g.Group(children),
	)
}