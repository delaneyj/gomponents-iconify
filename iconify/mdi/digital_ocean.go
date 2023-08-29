package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalOcean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 12H2C2 6.5 6.5 2 12 2s10 4.5 10 10s-4.5 10-10 10v-4H8v-4h4v4c3.32 0 6-2.69 6-6s-2.69-6-6-6s-6 2.69-6 6m2 6v3H5v-3h3m-5-2h2v2H3v-2Z"/>`),
		g.Group(children),
	)
}