package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symbolab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 25.38v5.732a4.246 4.246 0 0 1-4.246 4.246h0a4.233 4.233 0 0 1-3.003-1.243"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 18.374v7.006a4.246 4.246 0 0 1-4.246 4.246h0a4.246 4.246 0 0 1-4.246-4.246v-7.006m-15.172 9.391a4.75 4.75 0 0 0 4.165 1.861h2.514a4.242 4.242 0 0 0 4.237-4.246h0a4.242 4.242 0 0 0-4.237-4.246h-2.778A4.242 4.242 0 0 1 12 16.888h0a4.242 4.242 0 0 1 4.237-4.247h2.515a4.75 4.75 0 0 1 4.164 1.861"/>`),
		g.Group(children),
	)
}