package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 48v160a8 8 0 0 1-16 0v-72H43.31l18.35 18.34a8 8 0 0 1-11.32 11.32l-32-32a8 8 0 0 1 0-11.32l32-32a8 8 0 0 1 11.32 11.32L43.31 120H96V48a8 8 0 0 1 16 0Zm125.66 74.34l-32-32a8 8 0 0 0-11.32 11.32L212.69 120H160V48a8 8 0 0 0-16 0v160a8 8 0 0 0 16 0v-72h52.69l-18.35 18.34a8 8 0 0 0 11.32 11.32l32-32a8 8 0 0 0 0-11.32Z"/>`),
		g.Group(children),
	)
}