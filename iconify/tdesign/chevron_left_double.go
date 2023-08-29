package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.414 7.5l-4.5 4.5l4.5 4.5L17 17.914L11.086 12L17 6.086L18.414 7.5Zm-6.5 0l-4.5 4.5l4.5 4.5l-1.414 1.414L4.586 12L10.5 6.086L11.914 7.5Z"/>`),
		g.Group(children),
	)
}