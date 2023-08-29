package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoachLampVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 2L12 4l-2 2.31L5 9h2l2.5 9l2.5 2l.5 2h1l.5-2l2.5-2L19 9h2l-5-2.69L14 4l-.5-2M9 9h8l-2.22 8h-3.56M3 14v8h8.5l-.5-2H8l-3-3v-3Z"/>`),
		g.Group(children),
	)
}