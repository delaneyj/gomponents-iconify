package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatEmptyTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 0 0-4.386 7.403l-.592 1.947a.5.5 0 0 0 .624.624l1.945-.592A5 5 0 1 0 6 1Z"/>`),
		g.Group(children),
	)
}