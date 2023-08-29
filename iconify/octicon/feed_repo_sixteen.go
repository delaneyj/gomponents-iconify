package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedRepoSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16ZM5.5 4A1.5 1.5 0 0 0 4 5.5v5c0 .828.5 1.5 1 1.5v-1a1 1 0 0 1 1-1h5v1h-1v1h1.5a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5Zm.5 7.25v2.514a.25.25 0 0 0 .426.178l.898-.888a.25.25 0 0 1 .352 0l.898.888A.25.25 0 0 0 9 13.764V11H6.25a.25.25 0 0 0-.25.25Z"/>`),
		g.Group(children),
	)
}