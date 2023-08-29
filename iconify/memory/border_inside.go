package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h8V2h2v8h8v2h-8v8h-2v-8H2v-2m4 8h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}