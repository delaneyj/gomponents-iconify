package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutCardinalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M91.76 60.24a6 6 0 0 1 0-8.48l32-32a6 6 0 0 1 8.48 0l32 32a6 6 0 1 1-8.48 8.48L134 38.49V96a6 6 0 0 1-12 0V38.49l-21.76 21.75a6 6 0 0 1-8.48 0Zm64 135.52L134 217.51V160a6 6 0 0 0-12 0v57.51l-21.76-21.75a6 6 0 0 0-8.48 8.48l32 32a6 6 0 0 0 8.48 0l32-32a6 6 0 0 0-8.48-8.48Zm80.48-72l-32-32a6 6 0 0 0-8.48 8.48L217.51 122H160a6 6 0 0 0 0 12h57.51l-21.75 21.76a6 6 0 1 0 8.48 8.48l32-32a6 6 0 0 0 0-8.48ZM38.49 134H96a6 6 0 0 0 0-12H38.49l21.75-21.76a6 6 0 0 0-8.48-8.48l-32 32a6 6 0 0 0 0 8.48l32 32a6 6 0 0 0 8.48-8.48Z"/>`),
		g.Group(children),
	)
}