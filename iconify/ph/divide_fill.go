package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivideFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-80 32a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm0 128a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm56-56H72a8 8 0 0 1 0-16h112a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}