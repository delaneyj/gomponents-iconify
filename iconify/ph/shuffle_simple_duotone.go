package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m48 48l80 80l-80 80Zm80 80l80 80V48Z" opacity=".2"/><path d="M216 48v40a8 8 0 0 1-16 0V67.31L156.28 111A8 8 0 0 1 145 99.72L188.69 56H168a8 8 0 0 1 0-16h40a8 8 0 0 1 8 8Zm-8 112a8 8 0 0 0-8 8v20.69L53.66 42.34a8 8 0 0 0-11.32 11.32L188.69 200H168a8 8 0 0 0 0 16h40a8 8 0 0 0 8-8v-40a8 8 0 0 0-8-8ZM99.72 145l-57.38 57.34a8 8 0 0 0 11.32 11.32L111 156.28A8 8 0 0 0 99.72 145Z"/></g>`),
		g.Group(children),
	)
}