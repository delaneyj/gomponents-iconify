package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.342 2.447L6.236 6H23v16h-1v-1v1H1V6.382L12.447.658l.895 1.79ZM21 20V8H3v12h18ZM9 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm10-3h4v2h-4v-2Zm0 4h4v2h-4v-2Z"/>`),
		g.Group(children),
	)
}