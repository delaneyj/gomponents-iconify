package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6c.768 0 1.47-.289 2-.764c.53.475 1.232.764 2 .764h6a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-6c-.768 0-1.47.289-2 .764A2.989 2.989 0 0 0 10 4H4Zm7 3v10a1 1 0 0 1-1 1H4V6h6a1 1 0 0 1 1 1Zm2 10V7a1 1 0 0 1 1-1h6v12h-6a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}