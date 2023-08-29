package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 12v-.001c0-6.627-5.372-12-12-12c-6.627 0-12 5.372-12 12V12c0 6.627 5.372 12 12 12c6.627 0 12-5.372 12-12z"/>`),
		g.Group(children),
	)
}