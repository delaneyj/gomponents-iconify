package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRecordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8c2.205 0 4 1.794 4 4s-1.795 4-4 4s-4-1.794-4-4s1.795-4 4-4m0-2a6 6 0 1 0 0 12a6 6 0 0 0 0-12z"/>`),
		g.Group(children),
	)
}