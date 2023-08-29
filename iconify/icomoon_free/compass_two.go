package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zM1.5 8a6.5 6.5 0 0 1 10.93-4.756L6 6l-2.756 6.43A6.476 6.476 0 0 1 1.5 8zm7.643 1.143l-4.001 1.715l1.715-4.001l2.286 2.286zM8 14.5a6.476 6.476 0 0 1-4.43-1.744L10 10l2.756-6.43A6.5 6.5 0 0 1 8 14.5z"/>`),
		g.Group(children),
	)
}