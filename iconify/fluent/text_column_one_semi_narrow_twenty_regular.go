package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnOneSemiNarrowTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7Zm0 3a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7ZM6 11.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7Z"/>`),
		g.Group(children),
	)
}