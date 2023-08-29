package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanbanFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 56v96a8 8 0 0 1-8 8h-40a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h40a8 8 0 0 1 8 8Zm64-8h-40a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h40a8 8 0 0 0 8-8V56a8 8 0 0 0-8-8Zm0 80h-40a8 8 0 0 0-8 8v40a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-40a8 8 0 0 0-8-8ZM80 48H40a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h40a8 8 0 0 0 8-8V56a8 8 0 0 0-8-8Zm0 80H40a8 8 0 0 0-8 8v72a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-72a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}