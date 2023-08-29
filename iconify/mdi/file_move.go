package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17h4v-3l5 4.5l-5 4.5v-3h-4v-3m-1-8h5.5L13 3.5V9M6 2h8l6 6v4.34c-.63-.22-1.3-.34-2-.34a6 6 0 0 0-6 6c0 1.54.58 2.94 1.53 4H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}