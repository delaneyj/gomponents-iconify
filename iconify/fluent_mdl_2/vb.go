package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H128V128h1792zm-119 119H247v1554h1554V247zM639 1664L368 896h140l189 583q5 16 8 32t8 33q2-17 6-33t10-33l194-582h135l-276 768H639zm1025-221q0 56-23 97t-62 69t-86 41t-99 14h-242V896h240q41 0 83 9t77 30t56 54t22 83q0 67-37 114t-100 70q78 9 124 58t47 129zm-387-226h81q29 0 54-7t45-21t30-37t11-55q0-29-11-48t-29-30t-41-16t-49-5h-91v219zm107 345q30 0 56-7t47-21t31-38t12-57q0-36-14-59t-36-37t-52-19t-60-5h-91v243h107z"/>`),
		g.Group(children),
	)
}