package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.002 15.999a2 2 0 1 0-.004 4a2 2 0 0 0 .004-4zM6 4a2 2 0 0 0 0 4c5.514 0 10 4.486 10 10a2 2 0 0 0 4 0c0-7.72-6.28-14-14-14zm0 6a2 2 0 0 0 0 4c2.205 0 4 1.794 4 4a2 2 0 0 0 4 0c0-4.411-3.589-8-8-8z"/>`),
		g.Group(children),
	)
}