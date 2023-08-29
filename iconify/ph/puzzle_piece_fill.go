package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzlePieceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M165.78 224H208a16 16 0 0 0 16-16v-37.65a8 8 0 0 0-11.06-7.35a23.37 23.37 0 0 1-8.94 1.77c-13.23 0-24-11.1-24-24.73s10.77-24.73 24-24.73a23.37 23.37 0 0 1 8.94 1.77a8 8 0 0 0 11.06-7.43V72a16 16 0 0 0-16-16h-36.22a35.36 35.36 0 0 0 .22-4a36 36 0 0 0-72 0a35.36 35.36 0 0 0 .22 4H64a16 16 0 0 0-16 16v32.22a35.36 35.36 0 0 0-4-.22a36 36 0 0 0 0 72a35.36 35.36 0 0 0 4-.22V208a16 16 0 0 0 16 16h42.22"/>`),
		g.Group(children),
	)
}