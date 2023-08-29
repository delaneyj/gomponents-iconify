package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-36 136a12 12 0 0 0 12-12a8 8 0 0 1 16 0a28 28 0 0 1-56 0V96h-32v80a8 8 0 0 1-16 0V96h-8a24 24 0 0 0-24 24a8 8 0 0 1-16 0a40 40 0 0 1 40-40h104a8 8 0 0 1 0 16h-32v60a12 12 0 0 0 12 12Z"/>`),
		g.Group(children),
	)
}