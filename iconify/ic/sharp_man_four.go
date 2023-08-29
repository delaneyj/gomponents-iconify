package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpManFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.96 7L10 22h4l2.04-15z"/><circle cx="12" cy="4" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}