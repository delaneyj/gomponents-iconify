package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5v2H3v14h7v-8H5V4h14v6h-5v8h3v2h-6v2h8v-4h2V4h-2V2zm-3 10h3v4h-3v-4zm-8 0v4H5v-4h3z"/>`),
		g.Group(children),
	)
}