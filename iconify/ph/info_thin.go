package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 176a4 4 0 0 1-4 4a12 12 0 0 1-12-12v-40a4 4 0 0 0-4-4a4 4 0 0 1 0-8a12 12 0 0 1 12 12v40a4 4 0 0 0 4 4a4 4 0 0 1 4 4Zm-16-84a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm104 36A100 100 0 1 1 128 28a100.11 100.11 0 0 1 100 100Zm-8 0a92 92 0 1 0-92 92a92.1 92.1 0 0 0 92-92Z"/>`),
		g.Group(children),
	)
}