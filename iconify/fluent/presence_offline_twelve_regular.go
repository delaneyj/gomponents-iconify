package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceOfflineTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.03 3.97a.75.75 0 0 1 0 1.06L7.06 6l.97.97a.75.75 0 0 1-1.06 1.06L6 7.06l-.97.97a.75.75 0 0 1-1.06-1.06L4.94 6l-.97-.97a.75.75 0 0 1 1.06-1.06l.97.97l.97-.97a.75.75 0 0 1 1.06 0ZM0 6a6 6 0 1 1 12 0A6 6 0 0 1 0 6Zm6-4.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9Z"/>`),
		g.Group(children),
	)
}