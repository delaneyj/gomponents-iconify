package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-88 128a8 8 0 0 1-16 0v-32H72v32a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0v32h32V80a8 8 0 0 1 16 0Zm80 0h-8v16a8 8 0 0 1-16 0v-16h-32a8 8 0 0 1-7.49-10.81l24-64a8 8 0 0 1 15 5.62l-20 53.19H176v-24a8 8 0 0 1 16 0v24h8a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}