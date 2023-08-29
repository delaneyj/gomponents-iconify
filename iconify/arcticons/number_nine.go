package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.49 37.5c4.89-.26 8.57-4.65 8.57-9.55v-8.4a9.06 9.06 0 0 0-9.06-9h0a9.06 9.06 0 0 0-9.06 9.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.94 19.56A9.06 9.06 0 0 0 24 28.61h0a9.06 9.06 0 0 0 9.06-9.06M16.54 34.54c1.65 2.16 3.73 3 6.62 3h1.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}