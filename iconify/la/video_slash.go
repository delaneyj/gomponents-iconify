package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.719 2.281L2.28 3.72l26 26l1.438-1.438L24 22.563v-1.938l6 3V8.375l-6 3V8H9.437zM2 8v16h18l-2-2H4V10h2L4 8zm9.438 2H22v10.563zM28 11.625v8.75l-4-2v-4.75z"/>`),
		g.Group(children),
	)
}