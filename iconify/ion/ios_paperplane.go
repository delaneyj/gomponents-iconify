package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPaperplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 249.6l106 46.7L416 96z" fill="currentColor"/><path d="M416 96L217.9 311.7 269.8 416z" fill="currentColor"/>`),
		g.Group(children),
	)
}