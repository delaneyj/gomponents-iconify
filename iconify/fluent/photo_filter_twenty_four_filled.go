package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFilterTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 2a7.503 7.503 0 0 1 7.179 5.321a7.5 7.5 0 1 1-9.357 9.358A7.5 7.5 0 0 1 9.5 2Zm7.498 7.599L17 9.5A7.5 7.5 0 0 1 9.6 17a5.5 5.5 0 1 0 7.399-7.4ZM9.5 4a5.5 5.5 0 0 0-2.498 10.401L7 14.5A7.5 7.5 0 0 1 14.4 7a5.497 5.497 0 0 0-4.9-3Z"/>`),
		g.Group(children),
	)
}