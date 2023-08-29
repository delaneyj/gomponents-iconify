package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscendingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 128a8 8 0 0 1-8 8H48a8 8 0 0 1 0-16h72a8 8 0 0 1 8 8ZM48 72h136a8 8 0 0 0 0-16H48a8 8 0 0 0 0 16Zm56 112H48a8 8 0 0 0 0 16h56a8 8 0 0 0 0-16Zm127.39-19.06A8 8 0 0 0 224 160h-32v-48a8 8 0 0 0-16 0v48h-32a8 8 0 0 0-5.66 13.66l40 40a8 8 0 0 0 11.32 0l40-40a8 8 0 0 0 1.73-8.72Z"/>`),
		g.Group(children),
	)
}