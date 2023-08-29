package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignCenterRotateNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 4a.75.75 0 0 1 .75.75v14.5a.75.75 0 0 1-1.5 0V4.75a.75.75 0 0 1 .75-.75Zm-13 2a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-1.5 0V6.75A.75.75 0 0 1 5.25 6Zm7.25-3.25a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75Z"/>`),
		g.Group(children),
	)
}