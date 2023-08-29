package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleHangoutsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23a1 1 0 0 1-1-1v-1.477a9.842 9.842 0 1 1 10.705-8.527a12.531 12.531 0 0 1-9.466 10.975A.997.997 0 0 1 12 23Z" opacity=".5"/><path fill="currentColor" d="M9 11.689a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2Z"/><path fill="currentColor" d="M8.515 14.688a1 1 0 0 1 0-2a.501.501 0 0 0 .5-.5v-2.5a1 1 0 0 1 2 0v2.5a2.502 2.502 0 0 1-2.5 2.5zm6.485-3a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2z"/><path fill="currentColor" d="M14.515 14.688a1 1 0 0 1 0-2a.501.501 0 0 0 .5-.5v-2.5a1 1 0 0 1 2 0v2.5a2.502 2.502 0 0 1-2.5 2.5Z"/>`),
		g.Group(children),
	)
}