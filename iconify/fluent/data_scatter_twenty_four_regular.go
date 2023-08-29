package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScatterTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.75a.75.75 0 0 1 1.5 0V19.5h15.75a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75V3.75ZM17 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-1.5 3a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0ZM6 9a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm3-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm6 4.5a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-1.5 3a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}