package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BidLandscapeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-4.05V19h14v-8.75L13.05 17L9 12.95l-4 4Zm0-2.85l4-4l3.95 3.95L19 7.25V5H5v9.1Zm0-3.85v-3v6.8v-3.95v6.85v-4V17v-6.75Zm0 3.85V5v9.05v-3.95v4Zm0 2.85v-4V17v-6.75V19v-2.05Z"/>`),
		g.Group(children),
	)
}