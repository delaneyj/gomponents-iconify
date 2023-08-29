package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedPullRequestClosedSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.5 0a8 8 0 1 1 0 16a8 8 0 0 1 0-16ZM6 12.25a1.75 1.75 0 0 0 .5-3.428V6.928A1.752 1.752 0 0 0 6 3.5a1.75 1.75 0 0 0-.5 3.428v1.894A1.752 1.752 0 0 0 6 12.25Zm5-5a.5.5 0 0 0-.5.5v1.072a1.752 1.752 0 0 0 .5 3.428a1.75 1.75 0 0 0 .5-3.428V7.75a.5.5 0 0 0-.5-.5Zm1.255-2.763a.5.5 0 0 0-.707-.707l-.53.531l-.531-.531a.5.5 0 0 0-.707.707l.531.531l-.531.53a.5.5 0 0 0 .707.707l.531-.53l.53.53a.5.5 0 0 0 .707-.707l-.53-.53Z"/>`),
		g.Group(children),
	)
}