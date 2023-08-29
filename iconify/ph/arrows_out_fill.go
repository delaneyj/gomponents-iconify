package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M109.66 146.34a8 8 0 0 1 0 11.32L83.31 184l18.35 18.34A8 8 0 0 1 96 216H48a8 8 0 0 1-8-8v-48a8 8 0 0 1 13.66-5.66L72 172.69l26.34-26.35a8 8 0 0 1 11.32 0ZM83.31 72l18.35-18.34A8 8 0 0 0 96 40H48a8 8 0 0 0-8 8v48a8 8 0 0 0 13.66 5.66L72 83.31l26.34 26.35a8 8 0 0 0 11.32-11.32ZM208 40h-48a8 8 0 0 0-5.66 13.66L172.69 72l-26.35 26.34a8 8 0 0 0 11.32 11.32L184 83.31l18.34 18.35A8 8 0 0 0 216 96V48a8 8 0 0 0-8-8Zm3.06 112.61a8 8 0 0 0-8.72 1.73L184 172.69l-26.34-26.35a8 8 0 0 0-11.32 11.32L172.69 184l-18.35 18.34A8 8 0 0 0 160 216h48a8 8 0 0 0 8-8v-48a8 8 0 0 0-4.94-7.39Z"/>`),
		g.Group(children),
	)
}