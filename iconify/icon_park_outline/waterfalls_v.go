package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M23 20V6H6v14h17Zm19 22V28H25v14h17ZM31 6v14h11V6H31ZM6 28v14h11V28H6Z"/>`),
		g.Group(children),
	)
}