package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1063 288v60q0 285-532 653Q0 634 0 348v-60Q0 161 80.5 81T287 1q81 0 134 30.5T531 130q58-68 111-98.5T775 1q126 0 207 80t81 207zm-117 55v-50q0-58-39-113.5T812 118q-12-2-37-2q-53 0-83 19.5T620 204l-89 105l-89-105q-42-49-72-68.5T287 116q-24 0-37 2q-56 6-93.5 60T116 293v50q0 30 17 75q68 199 398 441q329-239 399-441q16-50 16-75z"/>`),
		g.Group(children),
	)
}