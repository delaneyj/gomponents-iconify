package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8 4l-6.32 6.32L3.35 12L8 7.35L12.65 12l1.67-1.68L8 4z"/>`),
		g.Group(children),
	)
}