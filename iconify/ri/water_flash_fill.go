package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterFlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.64 6.639L12.006.275l6.364 6.364a9 9 0 1 1-12.728 0Zm7.365 4.364v-4.5l-4.5 6.5h2.5v4.5l4.5-6.5h-2.5Z"/>`),
		g.Group(children),
	)
}