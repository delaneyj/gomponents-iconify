package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineDashesTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.785 2.223a.75.75 0 0 1 0 1.06l-1 1.001a.75.75 0 1 1-1.061-1.06l1-1.001a.75.75 0 0 1 1.061 0Zm-4.008 4a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 0 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0ZM13.28 10.72a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0Zm-4.503 5.563a.75.75 0 0 0-1.06-1.06l-1.5 1.5a.75.75 0 0 0 1.06 1.06l1.5-1.5ZM4.281 20.78a.75.75 0 1 0-1.06-1.06l-1.001 1a.75.75 0 1 0 1.06 1.061l1.001-1Z"/>`),
		g.Group(children),
	)
}