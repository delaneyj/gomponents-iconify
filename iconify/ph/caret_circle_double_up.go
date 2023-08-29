package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.54 54.46A104 104 0 0 0 54.46 201.54A104 104 0 0 0 201.54 54.46Zm-11.31 135.77a88 88 0 1 1 0-124.46a88.11 88.11 0 0 1 0 124.46Zm-24.57-27.89a8 8 0 0 1-11.32 11.32L128 147.31l-26.34 26.35a8 8 0 0 1-11.32-11.32l32-32a8 8 0 0 1 11.32 0Zm0-56a8 8 0 0 1-11.32 11.32L128 91.31l-26.34 26.35a8 8 0 0 1-11.32-11.32l32-32a8 8 0 0 1 11.32 0Z"/>`),
		g.Group(children),
	)
}