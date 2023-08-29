package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TflOyster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.593 42.5l13.244-13.245a2 2 0 0 0 0-2.828L22.012 12.602a2 2 0 0 0-2.828 0L5.5 26.286"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.233 35.86c-2.059-8.166-8.487-14.595-16.653-16.654"/>`),
		g.Group(children),
	)
}