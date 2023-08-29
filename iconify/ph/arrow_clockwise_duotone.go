package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 56v48h-48Z" opacity=".2"/><path d="M235.06 48.57a8 8 0 0 0-8.72 1.73L206.68 70l-10.93-10a96 96 0 1 0-69.07 164H128a95.44 95.44 0 0 0 65.88-26.19a8 8 0 0 0-11-11.63a80 80 0 1 1 1.68-114.78l.25.24l10.55 9.65l-17 17A8 8 0 0 0 184 112h48a8 8 0 0 0 8-8V56a8 8 0 0 0-4.94-7.43ZM224 96h-20.69L224 75.28Z"/></g>`),
		g.Group(children),
	)
}