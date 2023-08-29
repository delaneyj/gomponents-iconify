package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M117.66 138.34a8 8 0 0 1 0 11.32L83.31 184l18.35 18.34A8 8 0 0 1 96 216H48a8 8 0 0 1-8-8v-48a8 8 0 0 1 13.66-5.66L72 172.69l34.34-34.35a8 8 0 0 1 11.32 0ZM208 40h-48a8 8 0 0 0-5.66 13.66L172.69 72l-34.35 34.34a8 8 0 0 0 11.32 11.32L184 83.31l18.34 18.35A8 8 0 0 0 216 96V48a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}