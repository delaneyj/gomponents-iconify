package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 7.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm7-14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm7-21a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}