package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestClosedSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.25 1A2.25 2.25 0 0 1 4 5.372v5.256a2.251 2.251 0 1 1-1.5 0V5.372A2.251 2.251 0 0 1 3.25 1Zm9.5 5.5a.75.75 0 0 1 .75.75v3.378a2.251 2.251 0 1 1-1.5 0V7.25a.75.75 0 0 1 .75-.75Zm-2.03-5.273a.75.75 0 0 1 1.06 0l.97.97l.97-.97a.748.748 0 0 1 1.265.332a.75.75 0 0 1-.205.729l-.97.97l.97.97a.751.751 0 0 1-.018 1.042a.751.751 0 0 1-1.042.018l-.97-.97l-.97.97a.749.749 0 0 1-1.275-.326a.749.749 0 0 1 .215-.734l.97-.97l-.97-.97a.75.75 0 0 1 0-1.06ZM2.5 3.25a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0ZM3.25 12a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm9.5 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/>`),
		g.Group(children),
	)
}