package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3a1 1 0 0 0-2 0v3H3a1 1 0 0 0 0 2h3v9.75A4.25 4.25 0 0 0 10.25 22H20v3a1 1 0 1 0 2 0v-3h3a1 1 0 1 0 0-2H10.25A2.25 2.25 0 0 1 8 17.75V3Zm12 7.25V19h2v-8.75A4.25 4.25 0 0 0 17.75 6H9v2h8.75A2.25 2.25 0 0 1 20 10.25Z"/>`),
		g.Group(children),
	)
}