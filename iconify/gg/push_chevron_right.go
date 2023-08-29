package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.414 5L6 6.414l5.657 5.657L6 17.728l1.414 1.414l7.071-7.07L7.415 5Zm8.929 14V5h2v14h-2Z"/>`),
		g.Group(children),
	)
}