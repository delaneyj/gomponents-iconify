package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 36v8M14 15L24 4l10 11l-3 3l8 8l-5 1.158L42 36H6l8-8.842L9 26l8-8l-3-3Z"/>`),
		g.Group(children),
	)
}