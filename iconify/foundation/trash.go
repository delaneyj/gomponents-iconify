package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M75.834 33.388h-51.67a2.372 2.372 0 0 0-2.375 2.373v49.887a2.376 2.376 0 0 0 2.375 2.377h51.67a2.374 2.374 0 0 0 2.375-2.377V35.76a2.37 2.37 0 0 0-2.375-2.372zm3.17-16.036H59.402v-2.999a2.373 2.373 0 0 0-2.373-2.377H42.971a2.375 2.375 0 0 0-2.375 2.377v2.999h-19.6a2.372 2.372 0 0 0-2.375 2.373v6.932a2.372 2.372 0 0 0 2.375 2.373h58.008a2.37 2.37 0 0 0 2.375-2.373v-6.932a2.37 2.37 0 0 0-2.375-2.373z"/>`),
		g.Group(children),
	)
}