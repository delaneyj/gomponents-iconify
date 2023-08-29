package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-79.43 157.66l-5.46 5.45a44 44 0 0 1-62.22-62.22l24-24a44.08 44.08 0 0 1 55.56-5.48a8 8 0 0 1-8.9 13.3a28 28 0 0 0-35.35 3.49l-24 24a28 28 0 0 0 39.6 39.6l5.45-5.46a8 8 0 0 1 11.32 11.32Zm66.54-66.55l-24 24a44.08 44.08 0 0 1-55.56 5.48a8 8 0 0 1 8.9-13.3a28.06 28.06 0 0 0 35.35-3.49l24-24a28 28 0 0 0-39.6-39.6l-5.45 5.46a8 8 0 0 1-11.32-11.32l5.46-5.45a44 44 0 0 1 62.22 62.22Z"/>`),
		g.Group(children),
	)
}