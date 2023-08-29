package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainymood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.615 22.975v-8h2.6c1.5 0 2.7 1.2 2.7 2.7s-1.2 2.7-2.7 2.7h-2.6m2.729-.003l2.571 2.603m7.275 0l-2.6-8l-2.7 8m.9-2.7h3.5m2.875-5.3v8m1.975 0v-8l5.299 8v-8m7.275 0l-2.6 4l-2.7-4m2.7 8v-4m-23.446 14.05v-8l4 8l4-8v8m1.89-2.7c0 1.5 1.2 2.7 2.6 2.7c1.5 0 2.7-1.2 2.7-2.7v-2.7c0-1.5-1.2-2.7-2.7-2.7s-2.6 1.2-2.6 2.7v2.7Zm7.293 0c0 1.5 1.2 2.7 2.6 2.7c1.5 0 2.7-1.2 2.7-2.7v-2.7c0-1.5-1.2-2.7-2.7-2.7s-2.6 1.2-2.6 2.7v2.7Zm7.234 2.7v-8h1.3c2.2 0 4 1.8 4 4h0c0 2.2-1.8 4-4 4h-1.3Z"/>`),
		g.Group(children),
	)
}