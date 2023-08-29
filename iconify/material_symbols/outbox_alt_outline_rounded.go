package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutboxAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.2 12.9q.55-.275.55-.9t-.55-.9L7.45 6.725q-.5-.25-.975.038T6 7.625v8.75q0 .575.475.863t.975.037L16.2 12.9ZM8 15v-2l3-1l-3-1V9l6.5 3L8 15Zm-3 6q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}