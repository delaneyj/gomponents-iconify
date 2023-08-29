package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-24 128H72a8 8 0 0 1 0-16h112a8 8 0 0 1 0 16Zm0-48H72a8 8 0 0 1 0-16h112a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}