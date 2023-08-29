package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsplashLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.001 10v4h4v-4h7v11h-18V10h7Zm-2 2h-3v7h14v-7h-3L16 16H8v-4Zm8-9v6h-8V3h8Zm-2 2h-4v2h4V5Z"/>`),
		g.Group(children),
	)
}