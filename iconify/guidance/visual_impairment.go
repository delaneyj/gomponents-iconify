package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualImpairment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.278 14.278a3.222 3.222 0 0 0-4.556-4.556M6.584 6.584l10.832 10.832M4.204 9.204l9.992 9.992M0 12c1.276 0 2.5 1 3.5 2l2.606 2.947a8.079 8.079 0 0 0 11.788 0L20.5 14c1-1 2.224-2 3.5-2c-1.276 0-2.5-1-3.5-2l-2.606-2.947a8.079 8.079 0 0 0-11.788 0L3.5 10c-1 1-2.224 2-3.5 2Z"/>`),
		g.Group(children),
	)
}