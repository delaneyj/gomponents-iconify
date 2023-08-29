package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPpf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 6v2c3.2 0 6.36 1.18 8.57 3.15C15.64 13 16.83 15.5 17 18h1.97A14 12.5 0 0 0 5 6m17 15H2V3h2v16h18Z"/>`),
		g.Group(children),
	)
}