package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Galaxywearable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.35 19.23l-2.39 9.54l-2.39-9.54l-2.38 9.54l-2.39-9.54m24.01 5.6a2.39 2.39 0 0 1 2.39-2.38h0m-2.39 0v6.32m-9.17-1.2a2.39 2.39 0 0 1-2.08 1.2h0a2.39 2.39 0 0 1-2.38-2.38v-1.56a2.39 2.39 0 0 1 2.38-2.38h0A2.39 2.39 0 0 1 26 24.83v.78h-4.82m11.54.78a2.39 2.39 0 0 1-2.38 2.38h0A2.39 2.39 0 0 1 28 26.39v-1.56a2.39 2.39 0 0 1 2.39-2.38h0a2.39 2.39 0 0 1 2.38 2.38m-.05 3.94v-6.32"/>`),
		g.Group(children),
	)
}