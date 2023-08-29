package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwaySixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 7.52V4.5a1 1 0 1 0-2 0V8a1 1 0 0 0 .375.78l2.5 2a1 1 0 1 0 1.25-1.56L8.5 7.52ZM8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Z"/>`),
		g.Group(children),
	)
}