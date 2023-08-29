package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMaximum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 6h2v2H8zm4 0h2v2h-2zm8 0h2v2h-2zm4 0h2v2h-2zm4 0h2v2h-2z"/><path fill="currentColor" d="m27.989 28l-.027-.163C26.51 19.074 24.066 9.089 18 8.089V6h-2v2.085c-6.084.978-8.533 10.977-9.986 19.752L5.987 28H4V8h2V6H4V2H2v26a2 2 0 0 0 2 2h26v-2ZM8.015 28c2.024-12.108 4.96-18 8.973-18s6.949 5.892 8.972 18Z"/>`),
		g.Group(children),
	)
}