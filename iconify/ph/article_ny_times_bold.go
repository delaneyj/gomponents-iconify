package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleNyTimesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 124a12 12 0 0 1 12-12h76a12 12 0 0 1 0 24h-76a12 12 0 0 1-12-12Zm88 28h-76a12 12 0 0 0 0 24h76a12 12 0 0 0 0-24Zm0 40H72a12 12 0 0 0 0 24h168a12 12 0 0 0 0-24ZM80 180a60.07 60.07 0 0 1-60-60a59.42 59.42 0 0 1 3.32-19.59A32 32 0 0 1 44 44a12 12 0 0 1 5.83 1.51l68.6 38.11A8 8 0 0 0 116 68a12 12 0 0 1 0-24a32 32 0 0 1 0 64a12 12 0 0 1-5.83-1.51L71.51 85A36 36 0 0 0 68 153.94V124a12 12 0 0 1 24 0v29.91A36 36 0 0 0 114 132a12 12 0 0 1 22.63 8A60.1 60.1 0 0 1 80 180ZM36.49 78.76a60 60 0 0 1 9.33-8l-4.25-2.36A8 8 0 0 0 36 76a8.22 8.22 0 0 0 .49 2.76Z"/>`),
		g.Group(children),
	)
}