package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 120v72H80v-72Z" opacity=".2"/><path d="m213.66 82.34l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V88a8 8 0 0 0-2.34-5.66ZM200 216H56V40h92.69L200 91.31V216Zm-24-104H80a8 8 0 0 0-8 8v72a8 8 0 0 0 8 8h96a8 8 0 0 0 8-8v-72a8 8 0 0 0-8-8Zm-8 72h-16v-32a8 8 0 0 0-16 0v32h-16v-32a8 8 0 0 0-16 0v32H88v-56h80Z"/></g>`),
		g.Group(children),
	)
}