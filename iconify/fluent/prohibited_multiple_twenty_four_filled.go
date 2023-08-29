package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.75 19.5a8.75 8.75 0 1 0 0-17.5a8.75 8.75 0 0 0 0 17.5Zm0-2a6.72 6.72 0 0 1-4.014-1.322l9.442-9.442A6.75 6.75 0 0 1 10.75 17.5Zm4.014-12.178l-9.442 9.442a6.75 6.75 0 0 1 9.441-9.441ZM20.5 10.75a9.75 9.75 0 0 1-9.75 9.75a9.743 9.743 0 0 1-3.053-.488A8.75 8.75 0 0 0 20.012 7.696a9.74 9.74 0 0 1 .488 3.054Z"/>`),
		g.Group(children),
	)
}