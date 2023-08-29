package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWifiProTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-1.925 0-3.1-1t-1.825-2.5q-.65-1.525-.863-3.25T3 12.025q0-2.55.388-4.413T4.8 4.476q1.025-1.25 2.75-1.862T12 2q2.725 0 4.45.613t2.75 1.862q1.025 1.275 1.413 3.138T21 12.024q0 1.5-.213 3.225t-.862 3.25Q19.275 20 18.1 21T15 22H9Zm7-9q.425 0 .713-.288T17 12q0-2.375-1.325-4.175T11.95 6q-.4 0-.675.3T11 7.025q0 .425.3.7t.725.275q1.425 0 2.188 1.238T15 12.024q0 .425.288.7T16 13Z"/>`),
		g.Group(children),
	)
}