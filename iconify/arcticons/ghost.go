package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.31 5.5h19a1.63 1.63 0 0 1 1.59 1.63v4.4a1.63 1.63 0 0 1-1.64 1.64h-19a1.63 1.63 0 0 1-1.64-1.64v-4.4A1.63 1.63 0 0 1 7.31 5.5Zm29.15 0h4.23a1.63 1.63 0 0 1 1.64 1.63v4.4a1.63 1.63 0 0 1-1.64 1.64h-4.23a1.63 1.63 0 0 1-1.64-1.64v-4.4a1.63 1.63 0 0 1 1.64-1.63ZM7.31 20.16h33.38a1.63 1.63 0 0 1 1.64 1.64v4.4a1.64 1.64 0 0 1-1.64 1.64H7.31a1.63 1.63 0 0 1-1.64-1.64v-4.4a1.63 1.63 0 0 1 1.64-1.64Zm0 14.66h11.94a1.64 1.64 0 0 1 1.64 1.64v4.41a1.63 1.63 0 0 1-1.64 1.63H7.31a1.62 1.62 0 0 1-1.64-1.63v-4.41a1.64 1.64 0 0 1 1.64-1.64Zm21.44 0h11.94a1.64 1.64 0 0 1 1.64 1.64v4.41a1.63 1.63 0 0 1-1.64 1.63H28.75a1.62 1.62 0 0 1-1.64-1.63v-4.41a1.64 1.64 0 0 1 1.64-1.64Zm0 0"/>`),
		g.Group(children),
	)
}