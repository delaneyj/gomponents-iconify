package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm4 2a2 2 0 0 0 0 4a2 2 0 0 0 0-4zm10 0a2 2 0 0 0 0 4a2 2 0 0 0 0-4zM11 19a2 2 0 0 0 0 4a2 2 0 0 0 0-4zm10 0a2 2 0 0 0 0 4a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}