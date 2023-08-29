package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.33 8.232L13.66 4l5 8.66l-3.33 4.232l-5.33.768L5 9l5.33-.768Zm-2.12 2.326l3.197-.46l1.998-2.54l2.846 4.93l-1.998 2.539l-3.198.46l-2.846-4.929Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}