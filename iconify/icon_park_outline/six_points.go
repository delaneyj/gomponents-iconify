package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixPoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24V12m0 12l-10.5 6.062L24 24Zm0 0l10.5 6.062L24 24Zm-10-8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm14 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm14-8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0-16a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM28 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}