package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2h2v20h-2v-2H8v-2h2V6H8V4h2V2zM6 8V6h2v2H6zm0 8H2V8h4v2H4v4h2v2zm0 0v2h2v-2H6zm13-5h3v2h-3v3h-2v-3h-3v-2h3V8h2v3z"/>`),
		g.Group(children),
	)
}