package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBookFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 112a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm64-72v176a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16v-24H32a8 8 0 0 1 0-16h16v-40H32a8 8 0 0 1 0-16h16V80H32a8 8 0 0 1 0-16h16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Zm-33.6 123.2a67.88 67.88 0 0 0-27.4-21.69a40 40 0 1 0-53.94 0A67.88 67.88 0 0 0 81.6 163.2a8 8 0 1 0 12.8 9.6a52 52 0 0 1 83.2 0a8 8 0 1 0 12.8-9.6Z"/>`),
		g.Group(children),
	)
}