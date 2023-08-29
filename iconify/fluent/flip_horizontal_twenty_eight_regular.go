package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.882 23.654a.75.75 0 0 1-.632.346h-9.5a.75.75 0 0 1-.75-.75V2.75a.75.75 0 0 1 1.43-.315l9.5 20.5a.75.75 0 0 1-.048.719ZM16.5 6.152V22.5h7.576L16.5 6.152ZM2.75 24a.75.75 0 0 1-.68-1.065l9.5-20.5A.75.75 0 0 1 13 2.75v20.5a.75.75 0 0 1-.75.75h-9.5Z"/>`),
		g.Group(children),
	)
}