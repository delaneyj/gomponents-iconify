package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPennantBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m243.94 92.67l-184-64A12 12 0 0 0 44 40v176a12 12 0 0 0 24 0v-39.47l175.94-61.2a12 12 0 0 0 0-22.66ZM68 151.12V56.88L203.47 104Z"/>`),
		g.Group(children),
	)
}