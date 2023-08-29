package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionFrontTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 4a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5H3.75ZM12 8.5A2.5 2.5 0 0 0 9.708 10H8.126a4.002 4.002 0 0 1 7.748 0h-1.582A2.5 2.5 0 0 0 12 8.5Zm-4 5h1.5v2.25a.75.75 0 0 1-1.5 0V13.5Zm8 0h-1.5v2.25a.75.75 0 0 0 1.5 0V13.5Zm4.25-2.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1 0-1.5h16.5ZM3 18.75a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}