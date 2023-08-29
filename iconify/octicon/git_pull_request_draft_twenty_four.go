package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestDraftTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.75 1.5a3.25 3.25 0 0 1 .745 6.414A.827.827 0 0 1 5.5 8v8a.827.827 0 0 1-.005.086A3.25 3.25 0 0 1 4.75 22.5a3.25 3.25 0 0 1-.745-6.414A.827.827 0 0 1 4 16V8c0-.029.002-.057.005-.086A3.25 3.25 0 0 1 4.75 1.5ZM16 19.25a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0ZM3 4.75a1.75 1.75 0 1 0 3.501-.001A1.75 1.75 0 0 0 3 4.75Zm0 14.5a1.75 1.75 0 1 0 3.501-.001A1.75 1.75 0 0 0 3 19.25Zm16.25-1.75a1.75 1.75 0 1 0 .001 3.501a1.75 1.75 0 0 0-.001-3.501Zm0-11.5a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5ZM21 11.25a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Z"/>`),
		g.Group(children),
	)
}