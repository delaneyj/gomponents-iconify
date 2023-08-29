package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 8a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm4-2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}