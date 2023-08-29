package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.58 54.43a104 104 0 1 0 0 147.14a104.17 104.17 0 0 0 0-147.14Zm-35.9 119.25a8 8 0 0 1-11.32 0L128 147.32l-26.35 26.36a8 8 0 1 1-11.32-11.32l32-32a8 8 0 0 1 11.32 0l32 32a8 8 0 0 1 .03 11.32Zm0-56a8 8 0 0 1-11.32 0L128 91.29l-26.35 26.36a8 8 0 1 1-11.32-11.32l32-32a8 8 0 0 1 11.32 0l32 32a8 8 0 0 1 .03 11.32Z"/>`),
		g.Group(children),
	)
}