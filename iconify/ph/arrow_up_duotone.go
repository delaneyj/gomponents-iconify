package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 112H56l72-72Z" opacity=".2"/><path d="m205.66 106.34l-72-72a8 8 0 0 0-11.32 0l-72 72A8 8 0 0 0 56 120h64v96a8 8 0 0 0 16 0v-96h64a8 8 0 0 0 5.66-13.66ZM75.31 104L128 51.31L180.69 104Z"/></g>`),
		g.Group(children),
	)
}