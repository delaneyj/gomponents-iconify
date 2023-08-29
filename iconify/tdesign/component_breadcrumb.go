package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentBreadcrumb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6.5v11h7.914l5.5-5.5l-5.5-5.5H2Zm2 2h5.086l3.5 3.5l-3.5 3.5H4v-7Zm10.586-1l4.5 4.5l-4.5 4.5L16 17.914L21.914 12L16 6.086L14.586 7.5Z"/>`),
		g.Group(children),
	)
}