package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceholderFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-20.69 155.31a8 8 0 0 1-11.31 0L68.69 80A8 8 0 0 1 80 68.69L187.31 176a8 8 0 0 1 0 11.31Z"/>`),
		g.Group(children),
	)
}