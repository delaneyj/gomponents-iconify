package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5ZM9.128 2A3.496 3.496 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2h6.628ZM13 7.5v6h-2v-6h2Zm-2 7.496h2.004V17H11v-2.004Z"/>`),
		g.Group(children),
	)
}