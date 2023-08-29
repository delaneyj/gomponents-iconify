package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1.34 0a.5.5 0 0 0 .16 1H2v1.41C1.91 2.58.8 4.72.34 5.5C.18 5.76 0 6.11 0 6.56c0 .39.15.77.41 1.03S1.05 8 1.44 8h5.13c.38 0 .74-.16 1-.41h.03c.26-.26.41-.64.41-1.03c0-.45-.19-.8-.34-1.06c-.46-.78-1.57-2.92-1.66-3.09V1h.5a.5.5 0 1 0 0-1h-5a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM3 1h2v1.63l.06.09S5.69 3.95 6.25 5h-4.5c.56-1.05 1.19-2.28 1.19-2.28L3 2.63V1z"/>`),
		g.Group(children),
	)
}