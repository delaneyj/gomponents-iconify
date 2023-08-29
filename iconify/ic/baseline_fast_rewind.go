package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineFastRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18V6l-8.5 6l8.5 6zm.5-6l8.5 6V6l-8.5 6z"/>`),
		g.Group(children),
	)
}