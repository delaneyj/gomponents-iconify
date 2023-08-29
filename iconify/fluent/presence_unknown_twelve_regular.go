package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceUnknownTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9ZM0 6a6 6 0 1 1 12 0A6 6 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}