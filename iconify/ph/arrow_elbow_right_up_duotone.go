package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 96h-96l48-48Z" opacity=".2"/><path d="m229.66 90.34l-48-48a8 8 0 0 0-11.32 0l-48 48A8 8 0 0 0 128 104h40v80H32a8 8 0 0 0 0 16h144a8 8 0 0 0 8-8v-88h40a8 8 0 0 0 5.66-13.66ZM147.31 88L176 59.31L204.69 88Z"/></g>`),
		g.Group(children),
	)
}