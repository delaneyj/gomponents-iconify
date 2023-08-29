package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanbanBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a12 12 0 0 0-12 12v152a20 20 0 0 0 20 20h40a20 20 0 0 0 20-20v-44h40v12a20 20 0 0 0 20 20h40a20 20 0 0 0 20-20V56a12 12 0 0 0-12-12Zm-12 64h-32V68h32ZM84 68v40H52V68Zm0 136H52v-72h32Zm24-64V68h40v72Zm64 32v-40h32v40Z"/>`),
		g.Group(children),
	)
}