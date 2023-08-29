package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightRotateNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 5a.75.75 0 0 1 .75.75v15.5a.75.75 0 0 1-1.5 0V5.75a.75.75 0 0 1 .75-.75Zm-13 5a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-1.5 0v-10.5a.75.75 0 0 1 .75-.75Zm7.25-7.25a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75Z"/>`),
		g.Group(children),
	)
}