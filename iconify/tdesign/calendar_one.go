package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1h2ZM4 6v3h16V6H4Zm16 5H4v9h16v-9ZM7 13h2.004v2.004H7V13Zm4 0h2.004v2.004H11V13Zm4 0h2.004v2.004H15V13Zm-8 3h2.004v2.004H7V16Zm4 0h2.004v2.004H11V16Zm4 0h2.004v2.004H15V16Z"/>`),
		g.Group(children),
	)
}