package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaEject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16H7a2 2 0 0 0 0 4h10a2 2 0 0 0 0-4zm1.433-5.396L12 4l-6.433 6.604A2 2 0 0 0 7 14h10a2 2 0 0 0 1.433-3.396z"/>`),
		g.Group(children),
	)
}