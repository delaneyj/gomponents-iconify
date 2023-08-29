package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenVariantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L16.11 18H5V9H4c-.55 0-1-.45-1-1V7c0-.55.45-1 1-1h.11l-3-3l1.28-1.27l19.72 19.73l-1.27 1.27M19 9h1c.55 0 1-.45 1-1V7c0-.55-.45-1-1-1H9.2l9.8 9.8V9Z"/>`),
		g.Group(children),
	)
}