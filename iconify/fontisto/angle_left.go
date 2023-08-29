package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.2 2.43L16.778 0L4.8 12l11.978 12l2.422-2.43L9.653 12z"/>`),
		g.Group(children),
	)
}