package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M195.88 195.88a96 96 0 1 1 0-135.76a96 96 0 0 1 0 135.76Z" opacity=".2"/><path d="M201.54 54.46A104 104 0 0 0 54.46 201.54A104 104 0 0 0 201.54 54.46Zm-11.31 135.77a88 88 0 1 1 0-124.46a88.11 88.11 0 0 1 0 124.46ZM165.66 82.34a8 8 0 0 1 0 11.32l-32 32a8 8 0 0 1-11.32 0l-32-32a8 8 0 0 1 11.32-11.32L128 108.69l26.34-26.35a8 8 0 0 1 11.32 0Zm0 56a8 8 0 0 1 0 11.32l-32 32a8 8 0 0 1-11.32 0l-32-32a8 8 0 0 1 11.32-11.32L128 164.69l26.34-26.35a8 8 0 0 1 11.32 0Z"/></g>`),
		g.Group(children),
	)
}