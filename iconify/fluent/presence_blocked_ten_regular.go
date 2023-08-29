package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceBlockedTenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 5A5 5 0 1 0 0 5a5 5 0 0 0 10 0ZM9 5a4 4 0 0 1-6.453 3.16L8.16 2.547C8.686 3.224 9 4.076 9 5ZM7.453 1.84L1.84 7.453A4 4 0 0 1 7.453 1.84Z"/>`),
		g.Group(children),
	)
}