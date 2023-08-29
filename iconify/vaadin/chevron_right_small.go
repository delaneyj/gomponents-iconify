package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 8L5.68 1.68L4 3.35L8.65 8L4 12.65l1.68 1.67L12 8z"/>`),
		g.Group(children),
	)
}