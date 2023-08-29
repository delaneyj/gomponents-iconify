package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screencast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.35 4.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Zm4.36 8.1h14.62a1.61 1.61 0 0 1 1.61 1.61V16l4-2.29v9.16l-4-2.29v1.79A1.61 1.61 0 0 1 29.33 24H14.71a1.61 1.61 0 0 1-1.62-1.62v-8.17a1.61 1.61 0 0 1 1.62-1.61Zm-.31 19.62h2.73a1 1 0 0 1 1 1v2.74a1 1 0 0 1-1 1H14.4a1 1 0 0 1-1-1v-2.77a1 1 0 0 1 1-.97Zm8.23 0h2.74a1 1 0 0 1 1 1v2.74a1 1 0 0 1-1 1h-2.74a1 1 0 0 1-1-1v-2.77a1 1 0 0 1 1-.97Zm8.24 0h2.73a1 1 0 0 1 1 1v2.74a1 1 0 0 1-1 1h-2.73a1 1 0 0 1-1-1v-2.77a1 1 0 0 1 1-.97Z"/>`),
		g.Group(children),
	)
}