package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanbanDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 56v64h-56V56ZM40 208a8 8 0 0 0 8 8h40a8 8 0 0 0 8-8v-88H40Z" opacity=".2"/><path d="M216 48H40a8 8 0 0 0-8 8v152a16 16 0 0 0 16 16h40a16 16 0 0 0 16-16v-48h48v16a16 16 0 0 0 16 16h40a16 16 0 0 0 16-16V56a8 8 0 0 0-8-8Zm-8 64h-40V64h40ZM88 64v48H48V64Zm0 144H48v-80h40Zm16-64V64h48v80Zm64 32v-48h40v48Z"/></g>`),
		g.Group(children),
	)
}