package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Koho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.1 20v8m4.3 0l-3.3-4l3.3-4m-3.3 4h-1m7.5 1.4a2.65 2.65 0 1 0 5.3 0v-2.7a2.689 2.689 0 0 0-2.7-2.7a2.606 2.606 0 0 0-2.6 2.7Zm17 0a2.65 2.65 0 1 0 5.3 0v-2.7a2.689 2.689 0 0 0-2.7-2.7a2.606 2.606 0 0 0-2.6 2.7ZM25.1 20v8m5.3-8v8m-5.3-4h5.3M19.25 5.636V42.39"/>`),
		g.Group(children),
	)
}