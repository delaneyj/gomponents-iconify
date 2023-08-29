package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceDndTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 0a6 6 0 1 0 0 12A6 6 0 0 0 6 0ZM1.5 6a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0ZM3 6a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 3 6Z"/>`),
		g.Group(children),
	)
}