package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceOofTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.281 4.528a.75.75 0 0 0-1.06-1.06L3.218 5.47a.75.75 0 0 0 0 1.06l2.003 2.003a.75.75 0 0 0 1.06-1.061L5.56 6.75h2.69a.75.75 0 1 0 0-1.5H5.56l.722-.722ZM6 0a6 6 0 1 0 0 12A6 6 0 0 0 6 0ZM1.5 6a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z"/>`),
		g.Group(children),
	)
}