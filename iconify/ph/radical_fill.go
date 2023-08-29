package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadicalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm0 80a8 8 0 0 1-16 0v-8h-66.58l-30 75a8 8 0 0 1-14.86 0l-32-80a8 8 0 1 1 14.87-6L88 154.46L112.57 93a8 8 0 0 1 7.43-5h80a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}