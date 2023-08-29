package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionChristianityReligionJesusChristianityChristFishCulture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 4s-3.67 6-8 6a5.6 5.6 0 0 1-5-3a5.6 5.6 0 0 1 5-3c4.33 0 8 6 8 6"/>`),
		g.Group(children),
	)
}