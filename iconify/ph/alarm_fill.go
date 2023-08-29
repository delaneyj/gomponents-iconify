package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m61.66 29.66l-32 32a8 8 0 0 1-11.32-11.32l32-32a8 8 0 1 1 11.32 11.32Zm176 20.68l-32-32a8 8 0 0 0-11.32 11.32l32 32a8 8 0 0 0 11.32-11.32ZM224 128a96 96 0 1 1-96-96a96.11 96.11 0 0 1 96 96Zm-32 0a8 8 0 0 0-8-8h-48V72a8 8 0 0 0-16 0v56a8 8 0 0 0 8 8h56a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}