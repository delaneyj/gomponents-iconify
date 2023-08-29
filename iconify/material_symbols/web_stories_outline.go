package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebStoriesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20V4q.825 0 1.413.588T19 6v12q0 .825-.588 1.413T17 20ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h9q.825 0 1.413.588T15 4v16q0 .825-.588 1.413T13 22H4Zm17-4V6q.625 0 1.063.438T22.5 7.5v9q0 .625-.438 1.063T21 18ZM4 20h9V4H4v16ZM4 4v16V4Z"/>`),
		g.Group(children),
	)
}