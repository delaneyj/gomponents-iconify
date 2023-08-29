package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DietPlanNotebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1668 129h124v1919H256V129h126V0h128v129h258V0h128v129h258V0h128v129h258V0h128v129zm-4 1791V257H384v1663h1280zM1408 513v128H640V513h768zM640 1666v-128h768v128H640zm0-513v-128h768v128H640z"/>`),
		g.Group(children),
	)
}