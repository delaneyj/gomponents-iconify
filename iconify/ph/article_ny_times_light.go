package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleNyTimesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M154 104a6 6 0 0 1 6-6h80a6 6 0 0 1 0 12h-80a6 6 0 0 1-6-6Zm86 26h-80a6 6 0 0 0 0 12h80a6 6 0 0 0 0-12Zm0 32h-80a6 6 0 0 0 0 12h80a6 6 0 0 0 0-12Zm0 32H72a6 6 0 0 0 0 12h168a6 6 0 0 0 0-12ZM80 174a53.94 53.94 0 0 1-49.4-75.74A26 26 0 0 1 44 50a6 6 0 0 1 2.91.75l70.52 39.18A14 14 0 0 0 116 62a6 6 0 0 1 0-12a26 26 0 0 1 0 52a6 6 0 0 1-2.91-.75L72.46 78.67A42 42 0 0 0 74 161.56V120a6 6 0 0 1 12 0v41.56A42.06 42.06 0 0 0 119.61 134a6 6 0 0 1 11.32 4A54.11 54.11 0 0 1 80 174ZM36.64 87.9a54.29 54.29 0 0 1 21.41-17.23l-15.48-8.6a14 14 0 0 0-5.93 25.83Z"/>`),
		g.Group(children),
	)
}