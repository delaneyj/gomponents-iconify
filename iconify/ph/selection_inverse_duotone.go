package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionInverseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 48v160a8 8 0 0 1-2.34 5.66L42.34 42.34A8 8 0 0 1 48 40h160a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M152 216a8 8 0 0 1-8 8h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8ZM40 152a8 8 0 0 0 8-8v-32a8 8 0 0 0-16 0v32a8 8 0 0 0 8 8Zm32 56H48v-24a8 8 0 0 0-16 0v24a16 16 0 0 0 16 16h24a8 8 0 0 0 0-16ZM224 48v160a16 16 0 0 1-16 16h-24a8 8 0 0 1 0-16h12.69L48 59.31V72a8 8 0 0 1-16 0V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-16 0H59.31L208 196.69Z"/></g>`),
		g.Group(children),
	)
}