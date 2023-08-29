package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libremchat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.08 27.3h9.75a9.22 9.22 0 0 0 0-18.44H13.51a9.22 9.22 0 0 0-1.38 18.33v4.38Zm9.2-10.58a1.6 1.6 0 0 1 0 3.19a1.22 1.22 0 0 1-.27 0a1.6 1.6 0 0 1 .27-3.19Zm-7.49 1.58a1.6 1.6 0 1 1 1.59 1.61a1.6 1.6 0 0 1-1.59-1.61Zm-4.3-1.58a1.6 1.6 0 0 1 0 3.19a1.13 1.13 0 0 1-.26 0a1.6 1.6 0 0 1 .26-3.19Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.79 30.17a8.77 8.77 0 0 0 7.56 4.38h9.86l4.39 3.73v-3.84a8.76 8.76 0 0 0 7.49-6.59c1.17-4.7-2.14-10-7-10.63"/>`),
		g.Group(children),
	)
}