package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KebabHorizontalSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM1.5 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm13 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}