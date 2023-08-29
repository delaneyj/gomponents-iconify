package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 120H56l72-72Z" opacity=".2"/><path d="m205.66 114.34l-72-72a8 8 0 0 0-11.32 0l-72 72A8 8 0 0 0 56 128h144a8 8 0 0 0 5.66-13.66ZM75.31 112L128 59.31L180.69 112Z"/></g>`),
		g.Group(children),
	)
}