package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaultFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v136a16 16 0 0 0 16 16h16v16a8 8 0 0 0 16 0v-16h112v16a8 8 0 0 0 16 0v-16h16a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-8 96h-28.91a36 36 0 1 1 0-16H208a8 8 0 0 1 0 16Zm-44-8a20 20 0 1 1-20-20a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}