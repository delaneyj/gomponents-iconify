package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M896 1792q277 0 516-113.5t380-334.5q-80-50-240-45t-298 65q-90 39-159 98t-96 113v-449q201-28 332-153t131-326V81q0-34-23.5-57.5T1381 0q-21 10-67.5 56.5t-103 109T1139 243q-13-13-71-80.5t-105.5-115T896 0q-20 0-67.5 47.5t-105 115T653 243q-15-15-71.5-77.5t-103-109T410 0q-33 0-57 23.5T329 81v566q0 102 36.5 187.5T466 979t147 96.5t179 50.5v449q-26-53-95-112.5T538 1364q-138-60-298-65T0 1344q142 221 380.5 334.5T896 1792z"/>`),
		g.Group(children),
	)
}