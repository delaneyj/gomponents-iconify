package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3H5v4h2V3zm4 0H9v4h2V3zm2 0h2v4h-2V3zm8 6H3v12h14v-5h4V9zm-2 5h-2v-3h2v3zM5 11h10v8H5v-8z"/>`),
		g.Group(children),
	)
}