package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleNyTimesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 104a4 4 0 0 1 4-4h80a4 4 0 0 1 0 8h-80a4 4 0 0 1-4-4Zm84 28h-80a4 4 0 0 0 0 8h80a4 4 0 0 0 0-8Zm0 32h-80a4 4 0 0 0 0 8h80a4 4 0 0 0 0-8Zm0 32H72a4 4 0 0 0 0 8h168a4 4 0 0 0 0-8ZM80 172a51.94 51.94 0 0 1-46.8-74.59A24 24 0 0 1 44 52a4 4 0 0 1 1.94.5L117 92a16 16 0 0 0-1-32a4 4 0 0 1 0-8a24 24 0 0 1 0 48a4 4 0 0 1-1.94-.5L72.83 76.59A44 44 0 0 0 76 163.81V120a4 4 0 0 1 8 0v43.81a44.1 44.1 0 0 0 37.5-29.14a4 4 0 0 1 7.54 2.66A52.09 52.09 0 0 1 80 172ZM37.24 90.5A52.1 52.1 0 0 1 62.7 71L43 60a16 16 0 0 0-5.76 30.5Z"/>`),
		g.Group(children),
	)
}