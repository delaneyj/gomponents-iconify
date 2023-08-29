package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M417.059 497.692H272.891L49.574 250.448a24 24 0 0 1 2.007-34.148a90.409 90.409 0 0 1 129.507 10.789l18.494 22.6H224v-180a52 52 0 0 1 104 0v111.842l126.423 35.118A24.071 24.071 0 0 1 472 239.773v159.7a24 24 0 0 1-3.421 12.349Zm-129.95-32h111.832L440 397.26V245.854l-144-40V69.692a20 20 0 0 0-40 0v212h-71.582l-28.1-34.34a58.437 58.437 0 0 0-77.18-11.91Zm158.718-218.22l.033.009Z"/>`),
		g.Group(children),
	)
}