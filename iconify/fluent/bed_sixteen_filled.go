package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 6h1v-.5A.5.5 0 0 1 9 5h1.5a.5.5 0 0 1 .5.5V6h2V4a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2h2v-.5a.5.5 0 0 1 .5-.5H7a.5.5 0 0 1 .5.5V6Zm-4 1A2.5 2.5 0 0 0 1 9.5v4a.5.5 0 0 0 1 0V11h12v2.5a.5.5 0 0 0 1 0v-4A2.5 2.5 0 0 0 12.5 7h-9Z"/>`),
		g.Group(children),
	)
}