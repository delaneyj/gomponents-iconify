package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolboxLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 66h-50V56a22 22 0 0 0-22-22h-48a22 22 0 0 0-22 22v10H32a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h192a14 14 0 0 0 14-14V80a14 14 0 0 0-14-14ZM94 56a10 10 0 0 1 10-10h48a10 10 0 0 1 10 10v10H94ZM32 78h192a2 2 0 0 1 2 2v34h-36v-10a6 6 0 0 0-12 0v10H78v-10a6 6 0 0 0-12 0v10H30V80a2 2 0 0 1 2-2Zm192 116H32a2 2 0 0 1-2-2v-66h36v10a6 6 0 0 0 12 0v-10h100v10a6 6 0 0 0 12 0v-10h36v66a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}