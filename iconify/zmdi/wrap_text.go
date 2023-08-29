package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 341v-42h128v42H0zM341 43v42H0V43h341zm-64 128q36 0 61 25t25 60t-25 60t-61 25h-42v43l-64-64l64-64v43h48q17 0 29.5-12.5T325 256t-12.5-30.5T283 213H0v-42h277z"/>`),
		g.Group(children),
	)
}