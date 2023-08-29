package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CanWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.168 5.069c-.506-2.648-2.262-4.037-4.626-4.037c-2.362 0-4.125 1.384-4.636 4.027c-.473.012-.849.314-.849.69v1.847L3.904 5.575L3.9 5.573c.314-.551.385-1.075.121-1.346c-.39-.398-1.344-.099-2.131.668c-.787.77-1.111 1.715-.723 2.113c.266.273.797.214 1.355-.092l3.534 6.101v2.257c0 .382.39.692.873.692h9.158c.483 0 .873-.311.873-.692V5.749c.002-.359-.348-.646-.792-.68zM11.528 2c1.784 0 3.062.938 3.472 3H8c.41-2.062 1.744-3 3.527-3z"/>`),
		g.Group(children),
	)
}