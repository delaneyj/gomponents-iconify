package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscendingDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 168l-40 40l-40-40Z" opacity=".2"/><path d="M128 128a8 8 0 0 1-8 8H48a8 8 0 0 1 0-16h72a8 8 0 0 1 8 8ZM48 72h136a8 8 0 0 0 0-16H48a8 8 0 0 0 0 16Zm56 112H48a8 8 0 0 0 0 16h56a8 8 0 0 0 0-16Zm125.66-10.34l-40 40a8 8 0 0 1-11.32 0l-40-40A8 8 0 0 1 144 160h32v-48a8 8 0 0 1 16 0v48h32a8 8 0 0 1 5.66 13.66Zm-25 2.34h-41.35L184 196.69Z"/></g>`),
		g.Group(children),
	)
}