package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a2 2 0 1 1 4 0v11a2 2 0 1 1-4 0V6Zm16 0a2 2 0 1 1 4 0v11a2 2 0 1 1-4 0V6Zm-1.5 0h-9v11h9V6Z"/>`),
		g.Group(children),
	)
}