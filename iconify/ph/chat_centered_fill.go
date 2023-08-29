package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCenteredFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 56v128a16 16 0 0 1-16 16h-59.47l-14.84 24.29a16 16 0 0 1-27.41-.06L99.5 200.06L40 200a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}