package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighHeel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m231 164.19l-51-11.49L69.66 42.34a8 8 0 0 0-11.56.26C36.11 66.64 24 97 24 128v72a16 16 0 0 0 16 16h32a16 16 0 0 0 16-16v-48.94c2.49 1.45 4.94 3 7.34 4.64a112.45 112.45 0 0 1 40.55 50.39a15.9 15.9 0 0 0 14.83 9.91H240a16 16 0 0 0 16-16v-4.73a31.72 31.72 0 0 0-25-31.08ZM72 200H40v-63.71a110.88 110.88 0 0 1 32 7.12Zm168 0h-89.32a128.36 128.36 0 0 0-46.27-57.46a126.9 126.9 0 0 0-64.12-22.26a110.67 110.67 0 0 1 24.17-60.5l105.88 105.88a8 8 0 0 0 3.9 2.14l53.24 12A15.81 15.81 0 0 1 240 195.31Z"/>`),
		g.Group(children),
	)
}