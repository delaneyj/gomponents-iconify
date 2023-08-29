package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 18.365L50.14 136L256 253.635L461.86 136L256 18.365zm-154 168L50.14 216L256 333.635L461.86 216L410 186.365l-154 88l-154-88zm0 80L50.14 296L256 413.635L461.86 296L410 266.365l-154 88l-154-88zm0 80L50.14 376L256 493.635L461.86 376L410 346.365l-154 88l-154-88z"/>`),
		g.Group(children),
	)
}