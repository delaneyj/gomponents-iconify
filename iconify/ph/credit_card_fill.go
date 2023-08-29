package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16Zm-88 128h-16a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm64 0h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 0 16ZM32 88V64h192v24Z"/>`),
		g.Group(children),
	)
}