package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M49.501 8.249L38.251 20.75h-5l-11.25 12.499H17l-11.25 12.5V2H2v60h60V8.249H49.501M27.626 56.375h-9.688V35.124h9.688v21.251m16.25 0h-9.688v-33.75h9.688v33.75m16.249 0h-9.687V10.124h9.687v46.251"/>`),
		g.Group(children),
	)
}