package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 96a8 8 0 0 1-8 8H24a8 8 0 0 1 0-16h48a8 8 0 0 1 8 8Zm72 24h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm32-48h48a8 8 0 0 0 0-16h-48a8 8 0 0 0 0 16ZM72 120H24a8 8 0 0 0-8 8v64a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8v-64a8 8 0 0 0-8-8Zm160-32h-48a8 8 0 0 0-8 8v96a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8V96a8 8 0 0 0-8-8Zm-80 64h-48a8 8 0 0 0-8 8v32a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8v-32a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}