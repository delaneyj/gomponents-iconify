package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wordreference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.813 31.5v-15h4.911a5.038 5.038 0 0 1 0 10.075h-4.91m4.91 0l4.91 4.921m-13.01-9.933L20.53 31.5l-3.093-9.937l-3.094 9.937l-3.094-9.937M26.813 16.5h-2.146"/>`),
		g.Group(children),
	)
}