package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoStencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M59 60H13a1.002 1.002 0 0 1-1-1V13a1.002 1.002 0 0 1 1-1h46c.55 0 .998.447 1 1v46a.945.945 0 0 1-1 1z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M59 60H13a1.002 1.002 0 0 1-1-1V13a1.002 1.002 0 0 1 1-1h46c.55 0 .998.447 1 1v46a.945.945 0 0 1-1 1z"/><path d="M33.692 29.992a1 1 0 0 0-1 1v6.694l-2.11-2.694v3.243l2.323 2.966a1 1 0 0 0 1.787-.616v-9.593a1 1 0 0 0-.998-1h-.002zm-4.402 3.35l-2.323-2.966a1 1 0 0 0-1.787.616v9.593a1 1 0 1 0 2 0V33.89l2.11 2.695v-3.243zm12.389 6.148a2.39 2.39 0 0 1-1.78-2.332v-2.74a2.388 2.388 0 0 1 1.78-2.331v-2.032a4.36 4.36 0 0 0-3.78 4.363v2.74a4.36 4.36 0 0 0 3.78 4.363v-2.03zm1.292-9.435v2.032a2.388 2.388 0 0 1 1.78 2.33v2.74a2.495 2.495 0 0 1-1.78 2.32v2.04a4.483 4.483 0 0 0 3.78-4.358v-2.74a4.36 4.36 0 0 0-3.78-4.364z"/>`),
		g.Group(children),
	)
}