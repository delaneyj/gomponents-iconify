package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4.999a3.01 3.01 0 0 0-3.011 3l.005 9c0 2.209 1.793 4 4.002 4l9.003.001c1.655 0 3-1.346 3-3.001c.001-7.179-5.819-13-12.999-13zm1.001 14A1.998 1.998 0 0 1 7 17a1.999 1.999 0 0 1 2.001-2.001A1.996 1.996 0 0 1 11 17a1.995 1.995 0 0 1-1.999 1.999zm4.499 0a1.5 1.5 0 0 1-1.5-1.5c0-1.931-1.57-3.5-3.5-3.5a1.5 1.5 0 1 1 0-3c3.584 0 6.5 2.916 6.5 6.5a1.5 1.5 0 0 1-1.5 1.5zm4 0a1.5 1.5 0 0 1-1.5-1.5c0-4.136-3.364-7.5-7.5-7.5a1.5 1.5 0 1 1 0-3c5.79 0 10.5 4.71 10.5 10.5a1.5 1.5 0 0 1-1.5 1.5z"/>`),
		g.Group(children),
	)
}