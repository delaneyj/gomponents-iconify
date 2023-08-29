package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21Zm-1-4H6v1.5h12V17ZM6 15.5h12V14H6v1.5ZM6 12h12V6H6v6Zm0 5v1.5V17Zm0-1.5V14v1.5ZM6 12V6v6Zm0 2v-2v2Zm0 3v-1.5V17Z"/>`),
		g.Group(children),
	)
}