package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.586 16.5l4.5-4.5l-4.5-4.5L7 6.086L12.914 12L7 17.914L5.586 16.5Zm6.5 0l4.5-4.5l-4.5-4.5L13.5 6.086L19.414 12L13.5 17.914L12.086 16.5Z"/>`),
		g.Group(children),
	)
}