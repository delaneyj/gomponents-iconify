package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreakFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM96 64a8 8 0 0 1 16 0v16a8 8 0 0 1-16 0ZM64 96h16a8 8 0 0 1 0 16H64a8 8 0 0 1 0-16Zm64.08 85.66l-7.21 7.21a38 38 0 1 1-53.74-53.74l7.21-7.21a8 8 0 1 1 11.32 11.31l-7.22 7.21a22 22 0 0 0 31.12 31.12l7.21-7.22a8 8 0 1 1 11.31 11.32ZM160 192a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Zm32-32h-16a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm-3.13-39.13l-7.21 7.21a8 8 0 1 1-11.32-11.31l7.22-7.21a22 22 0 0 0-31.12-31.12l-7.21 7.22a8 8 0 1 1-11.31-11.32l7.21-7.21a38 38 0 1 1 53.74 53.74Z"/>`),
		g.Group(children),
	)
}