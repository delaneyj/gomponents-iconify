package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m160 56l-96 96V56Z" opacity=".2"/><path d="M224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8ZM56 152V56a8 8 0 0 1 8-8h96a8 8 0 0 1 5.66 13.66L123.31 104l58.35 58.34a8 8 0 0 1-11.32 11.32L112 115.31l-42.34 42.35A8 8 0 0 1 56 152Zm16-19.31l34.34-34.35L140.69 64H72Z"/></g>`),
		g.Group(children),
	)
}