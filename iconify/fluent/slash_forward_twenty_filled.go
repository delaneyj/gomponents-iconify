package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.016 2.049a.75.75 0 0 1 .435.967l-5.5 14.5a.75.75 0 1 1-1.402-.532l5.5-14.5a.75.75 0 0 1 .967-.435Z"/>`),
		g.Group(children),
	)
}