package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M110 48v160a6 6 0 0 1-12 0v-74H38.49l21.75 21.76a6 6 0 1 1-8.48 8.48l-32-32a6 6 0 0 1 0-8.48l32-32a6 6 0 0 1 8.48 8.48L38.49 122H98V48a6 6 0 0 1 12 0Zm126.24 75.76l-32-32a6 6 0 0 0-8.48 8.48L217.51 122H158V48a6 6 0 0 0-12 0v160a6 6 0 0 0 12 0v-74h59.51l-21.75 21.76a6 6 0 1 0 8.48 8.48l32-32a6 6 0 0 0 0-8.48Z"/>`),
		g.Group(children),
	)
}