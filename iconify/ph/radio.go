package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M104 176a8 8 0 0 1-8 8H64a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Zm-8-40H64a8 8 0 0 0 0 16h32a8 8 0 0 0 0-16Zm0-32H64a8 8 0 0 0 0 16h32a8 8 0 0 0 0-16Zm136-16v112a16 16 0 0 1-16 16H40a16 16 0 0 1-16-16V80a8 8 0 0 1 5.7-7.66l160-48a8 8 0 0 1 4.6 15.33L86.51 72H216a16 16 0 0 1 16 16Zm-16 112V88H40v112h176Zm-16-56a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm-16 0a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z"/>`),
		g.Group(children),
	)
}