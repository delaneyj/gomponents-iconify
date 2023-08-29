package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCounterClockwiseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 160v48a8 8 0 0 1-13.66 5.66l-16.06-16.07C174.13 214.9 151.36 224 128 224c-42.82 0-68.58-25.27-69.66-26.34a8 8 0 0 1 11.3-11.33C70 186.68 92.16 208 128 208c23.17 0 41.92-10.85 54.92-21.76l-20.58-20.58A8 8 0 0 1 168 152h48a8 8 0 0 1 8 8ZM95.39 99.06a8 8 0 0 0-1.73-8.72L73.08 69.76C86.08 58.85 104.83 48 128 48c35.95 0 58.13 21.44 58.36 21.68a8 8 0 0 0 11.3-11.34C196.58 57.27 170.82 32 128 32c-23.36 0-46.13 9.1-66.28 26.41L45.66 42.34A8 8 0 0 0 32 48v48a8 8 0 0 0 8 8h48a8 8 0 0 0 7.39-4.94Z"/>`),
		g.Group(children),
	)
}