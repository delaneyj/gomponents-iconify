package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColorAccentTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4.5 16.75a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-.75.75H5.25a.75.75 0 0 1-.75-.75v-3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}