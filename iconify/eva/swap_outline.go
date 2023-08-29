package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSwapOutline0"><g id="evaSwapOutline1"><path id="evaSwapOutline2" fill="currentColor" d="M4 9h13l-1.6 1.2a1 1 0 0 0-.2 1.4a1 1 0 0 0 .8.4a1 1 0 0 0 .6-.2l4-3a1 1 0 0 0 0-1.59l-3.86-3a1 1 0 0 0-1.23 1.58L17.08 7H4a1 1 0 0 0 0 2Zm16 7H7l1.6-1.2a1 1 0 0 0-1.2-1.6l-4 3a1 1 0 0 0 0 1.59l3.86 3a1 1 0 0 0 .61.21a1 1 0 0 0 .79-.39a1 1 0 0 0-.17-1.4L6.92 18H20a1 1 0 0 0 0-2Z"/></g></g>`),
		g.Group(children),
	)
}