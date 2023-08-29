package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopChromebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 15H4V5h16m-6 13h-4v-1h4m8 1V3H2v15H0v2h24v-2h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}