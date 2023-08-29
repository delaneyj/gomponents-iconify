package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19 8l-1.414 1.414L22.172 14H10a2 2 0 0 0-2 2v12h2V16h12.172l-4.586 4.586L19 22l7-7zM8 4h2v8H8z"/>`),
		g.Group(children),
	)
}