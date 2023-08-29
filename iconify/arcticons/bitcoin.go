package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M27.13 24.84a5.33 5.33 0 0 1-2.76 10.29l-8.49-2.28l5.51-20.58l8.49 2.28a5.33 5.33 0 1 1-2.75 10.29Zm0 0l-8.49-2.28m2.75-10.29l-3.01-.8m-2.5 21.38l-3.02-.81m11.15-19.06l.98-3.65m3.22 4.77l.98-3.65M17.52 37.2l.98-3.65m3.22 4.78l.98-3.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}