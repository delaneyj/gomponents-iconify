package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTextTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.004 7.75A3.75 3.75 0 0 1 5.754 4H22.25A3.75 3.75 0 0 1 26 7.75v12.5A3.75 3.75 0 0 1 22.25 24H5.755a3.75 3.75 0 0 1-3.75-3.75V7.75ZM7.728 9A.728.728 0 0 0 7 9.728v.044c0 .402.326.728.728.728h5.544A.728.728 0 0 0 14 9.772v-.044A.728.728 0 0 0 13.272 9H7.728ZM7 13.728v.044c0 .402.326.728.728.728h10.544a.728.728 0 0 0 .728-.728v-.044a.728.728 0 0 0-.728-.728H7.728a.728.728 0 0 0-.728.728ZM7.728 17a.728.728 0 0 0-.728.728v.044c0 .402.326.728.728.728h7.544a.728.728 0 0 0 .728-.728v-.044a.728.728 0 0 0-.728-.728H7.728Z"/>`),
		g.Group(children),
	)
}