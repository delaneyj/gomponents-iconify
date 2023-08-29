package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsClockwiseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M197.66 186.34a8 8 0 0 1 0 11.32C196.58 198.73 170.82 224 128 224c-23.36 0-46.13-9.1-66.28-26.41l-16.06 16.07A8 8 0 0 1 32 208v-48a8 8 0 0 1 8-8h48a8 8 0 0 1 5.66 13.66l-20.58 20.58C86.08 197.15 104.83 208 128 208c36.27 0 58.13-21.44 58.34-21.66a8 8 0 0 1 11.32 0Zm21.4-145.73a8 8 0 0 0-8.72 1.73l-16.06 16.07C174.13 41.1 151.36 32 128 32c-42.82 0-68.58 25.27-69.66 26.34a8 8 0 0 0 11.32 11.32C69.87 69.44 91.73 48 128 48c23.17 0 41.92 10.85 54.92 21.76l-20.58 20.58A8 8 0 0 0 168 104h48a8 8 0 0 0 8-8V48a8 8 0 0 0-4.94-7.39Z"/>`),
		g.Group(children),
	)
}