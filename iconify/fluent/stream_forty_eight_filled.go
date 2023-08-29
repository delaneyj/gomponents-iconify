package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M30.5 7C23.596 7 18 12.596 18 19.5v1a7.5 7.5 0 0 1-7.5 7.5h-5a1.5 1.5 0 0 1 0-3h5a4.5 4.5 0 0 0 4.5-4.5v-1C15 10.94 21.94 4 30.5 4h4.25a1.5 1.5 0 0 1 0 3H30.5Zm2 8a6.5 6.5 0 0 0-6.5 6.5v1C26 29.956 19.956 36 12.5 36h-7a1.5 1.5 0 0 1 0-3h7C18.299 33 23 28.299 23 22.5v-1a9.5 9.5 0 0 1 9.5-9.5h10a1.5 1.5 0 0 1 0 3h-10Zm1.5 9.364c0-.754.61-1.364 1.364-1.364H42.5a1.5 1.5 0 0 0 0-3h-7.136A4.364 4.364 0 0 0 31 24.364C31 33.552 23.552 41 14.364 41H13.5a1.5 1.5 0 0 0 0 3h.864C25.209 44 34 35.209 34 24.364Z"/>`),
		g.Group(children),
	)
}