package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 152a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm32-104v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-96 32a8 8 0 0 0-16 0v32H72V80a8 8 0 0 0-16 0v80a8 8 0 0 0 16 0v-32h40v32a8 8 0 0 0 16 0Zm80 72a32 32 0 0 0-32-32l11.55-20a8 8 0 0 0-13.86-8l-25.4 44l-.14.27A32 32 0 1 0 208 152Z"/>`),
		g.Group(children),
	)
}