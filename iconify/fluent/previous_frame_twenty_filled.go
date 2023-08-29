package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousFrameTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.5 3a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 1 0v-13a.5.5 0 0 0-.5-.5Zm-5.447.214A1.25 1.25 0 0 1 12 4.252v11.5a1.25 1.25 0 0 1-1.954 1.033l-8.499-5.793a1.25 1.25 0 0 1 .007-2.07l8.5-5.708Z"/>`),
		g.Group(children),
	)
}