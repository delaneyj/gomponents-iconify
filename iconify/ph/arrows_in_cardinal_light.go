package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInCardinalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M91.76 68.24a6 6 0 0 1 8.48-8.48L122 81.51V24a6 6 0 0 1 12 0v57.51l21.76-21.75a6 6 0 0 1 8.48 8.48l-32 32a6 6 0 0 1-8.48 0Zm40.48 87.52a6 6 0 0 0-8.48 0l-32 32a6 6 0 0 0 8.48 8.48L122 174.49V232a6 6 0 0 0 12 0v-57.51l21.76 21.75a6 6 0 0 0 8.48-8.48ZM232 122h-57.51l21.75-21.76a6 6 0 0 0-8.48-8.48l-32 32a6 6 0 0 0 0 8.48l32 32a6 6 0 0 0 8.48-8.48L174.49 134H232a6 6 0 0 0 0-12Zm-131.76 1.76l-32-32a6 6 0 0 0-8.48 8.48L81.51 122H24a6 6 0 0 0 0 12h57.51l-21.75 21.76a6 6 0 1 0 8.48 8.48l32-32a6 6 0 0 0 0-8.48Z"/>`),
		g.Group(children),
	)
}