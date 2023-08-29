package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.45 7.5H9.99A2 2 0 0 0 8.06 6h-1a2 2 0 0 0-1.93 1.5H2.55a.5.5 0 0 0-.5.5a.508.508 0 0 0 .5.5h2.58A2 2 0 0 0 7.06 10h1a2 2 0 0 0 1.93-1.5h11.46a.5.5 0 0 0 0-1ZM8.06 9h-1a1.006 1.006 0 0 1-1-.98V8a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2Zm13.39 6.5h-2.58a2 2 0 0 0-1.93-1.5h-1a2 2 0 0 0-1.93 1.5H2.55a.5.5 0 0 0 0 1h11.46a2 2 0 0 0 1.93 1.5h1a2 2 0 0 0 1.93-1.5h2.58a.5.5 0 0 0 .5-.5a.508.508 0 0 0-.5-.5Zm-3.51.5a1 1 0 0 1-1 1h-1a1 1 0 1 1 0-2h1a1.006 1.006 0 0 1 1 .98Z"/>`),
		g.Group(children),
	)
}