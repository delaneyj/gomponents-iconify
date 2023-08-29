package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCenteredDotsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h57.23l13.62 22.29a20 20 0 0 0 34.25.08L158.77 204H216a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 144h-55.47a20 20 0 0 0-17.1 9.63L128 208.33l-11.41-18.67A20.1 20.1 0 0 0 99.47 180H44V60h168ZM88 120a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm48 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}