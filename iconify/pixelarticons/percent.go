package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4h-2v2h-2v2h-2v2h-2v2h-2v2H8v2H6v2H4v2h2v-2h2v-2h2v-2h2v-2h2v-2h2V8h2V6h2V4zm-4 10h4v6h-6v-6h2zm2 4v-2h-2v2h2zM6 4h4v6H4V4h2zm2 4V6H6v2h2z"/>`),
		g.Group(children),
	)
}