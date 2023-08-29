package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.47 7.37v-.08l-.06-.15a.71.71 0 0 0-.07-.09a.94.94 0 0 0-.09-.12l-.09-.07l-.16-.08l-7.5-4.63a1 1 0 0 0-1.06 0L4 6.78l-.09.08l-.09.07a.94.94 0 0 0-.09.12a.71.71 0 0 0-.07.09l-.06.15v.08a1.15 1.15 0 0 0 0 .26v8.74a1 1 0 0 0 .47.85l7.5 4.63a.47.47 0 0 0 .15.06s.05 0 .08 0a.86.86 0 0 0 .52 0h.08a.47.47 0 0 0 .15-.06L20 17.22a1 1 0 0 0 .47-.85V7.63a1.15 1.15 0 0 0 0-.26ZM11 19.21l-5.5-3.4V9.43l5.5 3.39Zm1-8.12L6.4 7.63L12 4.18l5.6 3.45Zm6.5 4.72l-5.5 3.4v-6.39l5.5-3.39Z"/>`),
		g.Group(children),
	)
}