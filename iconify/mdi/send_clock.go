package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3v7l9 2l-9 2v7l7.27-3.11A7 7 0 0 0 16 23a7 7 0 0 0 7-7a7 7 0 0 0-7-7L2 3m14 8a5 5 0 0 1 5 5a5 5 0 0 1-5 5a5 5 0 0 1-5-5a5 5 0 0 1 5-5m-1 1.5v4l3 2l.75-1.25l-2.25-1.5V12.5H15Z"/>`),
		g.Group(children),
	)
}