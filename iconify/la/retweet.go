package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5 7l-5 5h4v13h17l-2-2H6V12h4zm4 0l2 2h13v11h-4l5 5l5-5h-4V7z"/>`),
		g.Group(children),
	)
}