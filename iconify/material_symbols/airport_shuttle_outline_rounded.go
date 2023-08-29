package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirportShuttleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19q-1.25 0-2.125-.875T3 16q-.825 0-1.413-.588T1 14V7q0-.825.588-1.413T3 5h13.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V14q0 .825-.588 1.413T21 16q0 1.25-.875 2.125T18 19q-1.25 0-2.125-.875T15 16H9q0 1.25-.875 2.125T6 19Zm9-9h4l-3-3h-1v3Zm-6 0h4V7H9v3Zm-6 0h4V7H3v3Zm3 7.25q.525 0 .888-.363T7.25 16q0-.525-.363-.888T6 14.75q-.525 0-.888.363T4.75 16q0 .525.363.888T6 17.25Zm12 0q.525 0 .888-.363T19.25 16q0-.525-.363-.888T18 14.75q-.525 0-.888.363T16.75 16q0 .525.363.888t.887.362ZM8.2 14h7.6q.425-.45.975-.725T18 13q.675 0 1.225.275T20.2 14h.8v-2H3v2h.8q.425-.45.975-.725T6 13q.675 0 1.225.275T8.2 14ZM21 12H3h18Z"/>`),
		g.Group(children),
	)
}