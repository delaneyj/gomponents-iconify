package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M1.5 32h22c.827 0 1.5-.673 1.5-1.5v-29c0-.827-.673-1.5-1.5-1.5h-22C.673 0 0 .673 0 1.5v29c0 .827.673 1.5 1.5 1.5zM1 1.5a.5.5 0 0 1 .5-.5h22a.5.5 0 0 1 .5.5v29a.5.5 0 0 1-.5.5h-22a.5.5 0 0 1-.5-.5v-29z"/><path d="M3.5 27h18a.5.5 0 0 0 .5-.5v-23a.5.5 0 0 0-.5-.5h-18a.5.5 0 0 0-.5.5v23a.5.5 0 0 0 .5.5zM4 4h17v22H4V4z"/><circle cx="13" cy="29" r="1"/></g>`),
		g.Group(children),
	)
}