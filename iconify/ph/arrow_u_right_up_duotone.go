package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 80h-96l48-48Z" opacity=".2"/><path d="m221.66 74.34l-48-48a8 8 0 0 0-11.32 0l-48 48A8 8 0 0 0 120 88h40v80a48 48 0 0 1-96 0V80a8 8 0 0 0-16 0v88a64 64 0 0 0 128 0V88h40a8 8 0 0 0 5.66-13.66ZM139.31 72L168 43.31L196.69 72Z"/></g>`),
		g.Group(children),
	)
}