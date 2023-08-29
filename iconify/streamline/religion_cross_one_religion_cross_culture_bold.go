package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionCrossOneReligionCrossCultureBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 3.5H9v-3H5v3H2v4h3v6h4v-6h3v-4z"/>`),
		g.Group(children),
	)
}