package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeriodicTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v2h2V4H2m18 0v2h2V4h-2M2 7v2h2V7H2m3 0v2h2V7H5m9 0v2h2V7h-2m3 0v2h2V7h-2m3 0v2h2V7h-2M2 10v2h2v-2H2m3 0v2h2v-2H5m3 0v2h2v-2H8m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2M2 13v2h2v-2H2m3 0v2h2v-2H5m3 0v2h2v-2H8m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2M5 17v2h2v-2H5m3 0v2h2v-2H8m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}