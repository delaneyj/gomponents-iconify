package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescendingDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 88h-80l40-40Z" opacity=".2"/><path d="M40 128a8 8 0 0 1 8-8h72a8 8 0 0 1 0 16H48a8 8 0 0 1-8-8Zm8-56h56a8 8 0 0 0 0-16H48a8 8 0 0 0 0 16Zm136 112H48a8 8 0 0 0 0 16h136a8 8 0 0 0 0-16Zm47.39-92.94A8 8 0 0 1 224 96h-32v48a8 8 0 0 1-16 0V96h-32a8 8 0 0 1-5.66-13.66l40-40a8 8 0 0 1 11.32 0l40 40a8 8 0 0 1 1.73 8.72ZM204.69 80L184 59.31L163.31 80Z"/></g>`),
		g.Group(children),
	)
}