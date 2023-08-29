package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceBusyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 10c0 5.523-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0s10 4.477 10 10Z"/>`),
		g.Group(children),
	)
}