package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 72v112a8 8 0 0 1-16 0V72a8 8 0 0 1 16 0Zm80-8h-48a8 8 0 0 0-8 8v112a8 8 0 0 0 16 0v-48h32a8 8 0 0 0 0-16h-32V80h40a8 8 0 0 0 0-16ZM96 120H72a8 8 0 0 0 0 16h16v16a24 24 0 0 1-48 0v-48a24 24 0 0 1 24-24c11.19 0 21.61 7.74 24.25 18a8 8 0 0 0 15.5-4C99.27 76.62 82.56 64 64 64a40 40 0 0 0-40 40v48a40 40 0 0 0 80 0v-24a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}