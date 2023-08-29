package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadsterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20H5v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-7.5l-1.243-.31A1 1 0 0 1 0 12.22v-.72a.5.5 0 0 1 .5-.5H2l2.48-5.788A2 2 0 0 1 6.32 4h11.36a2 2 0 0 1 1.838 1.212L22 11h1.5a.5.5 0 0 1 .5.5v.72a1 1 0 0 1-.758.97L22 13.5V21a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1Zm1-2v-5H4v5h16ZM5.477 11h13.046a1.001 1.001 0 0 0 .928-1.371L18 6H6L4.549 9.629A1 1 0 0 0 5.477 11ZM5 14c2.317 0 3.879.755 4.686 2.264a.5.5 0 0 1-.441.736H6a1 1 0 0 1-1-1v-2Zm14 0v2a1 1 0 0 1-1 1h-3.245a.5.5 0 0 1-.44-.736C15.12 14.754 16.682 14 19 14Z"/>`),
		g.Group(children),
	)
}