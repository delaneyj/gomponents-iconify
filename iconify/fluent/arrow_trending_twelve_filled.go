package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 1.5a.75.75 0 0 0 0 1.5h2.323L5.97 5.91l-.94-.94a.75.75 0 0 0-1.06 0L1.22 7.72a.75.75 0 0 0 1.06 1.06L4.5 6.56l.97.97a.75.75 0 0 0 1.089-.03L9.5 4.213V6.25a.75.75 0 0 0 1.5 0v-4a.75.75 0 0 0-.75-.75h-4Z"/>`),
		g.Group(children),
	)
}