package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedPullRequestDraftSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.5 8a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm7.25 2.5c0-.793-.527-1.462-1.25-1.678V6.928A1.752 1.752 0 0 0 6 3.5a1.75 1.75 0 0 0-.5 3.428v1.894A1.752 1.752 0 0 0 6 12.25a1.75 1.75 0 0 0 1.75-1.75ZM11 12.25a1.75 1.75 0 1 0 .001-3.499A1.75 1.75 0 0 0 11 12.25Zm0-4.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm.75-3.25a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Z"/>`),
		g.Group(children),
	)
}