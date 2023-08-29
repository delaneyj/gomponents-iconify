package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCenteredTextBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h57.23l13.62 22.29a20 20 0 0 0 34.25.08L158.77 204H216a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 144h-55.47a20 20 0 0 0-17.1 9.63L128 208.33l-11.41-18.67A20.1 20.1 0 0 0 99.47 180H44V60h168ZM84 100a12 12 0 0 1 12-12h64a12 12 0 0 1 0 24H96a12 12 0 0 1-12-12Zm0 40a12 12 0 0 1 12-12h64a12 12 0 0 1 0 24H96a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}