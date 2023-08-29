package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopArrowClockwiseSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4a2 2 0 0 1 2-2h3v1.75C7 4.44 7.56 5 8.25 5H14v7a2 2 0 0 1-2 2H9c.628-.836 1-1.874 1-3a4.98 4.98 0 0 0-.75-2.635V7.25a1.5 1.5 0 0 0-2.633-.983A4.994 4.994 0 0 0 5 6a4.978 4.978 0 0 0-3 1V4Zm6-2v1.75c0 .138.112.25.25.25H14a2 2 0 0 0-2-2H8Zm.247 5.198a.5.5 0 0 0-.997.052v.442A4 4 0 1 0 9 11a.5.5 0 0 0-1 0a3 3 0 1 1-1.341-2.5H6a.5.5 0 0 0 0 1h1.75a.5.5 0 0 0 .5-.5V7.25c0-.018 0-.035-.003-.052Z"/>`),
		g.Group(children),
	)
}