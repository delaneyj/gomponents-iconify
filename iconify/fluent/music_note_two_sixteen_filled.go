package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 5.21v5.058A2 2 0 1 0 13 12V1.926a.8.8 0 0 0-1.07-.754l-6.4 2.286a.8.8 0 0 0-.53.753v7.056A2 2 0 1 0 6 13V7.353l6-2.142Z"/>`),
		g.Group(children),
	)
}