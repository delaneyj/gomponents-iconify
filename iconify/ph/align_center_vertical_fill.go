package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 128a8 8 0 0 1-8 8h-16v48a16 16 0 0 1-16 16h-40a16 16 0 0 1-16-16v-48h-16v72a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16v-72H32a8 8 0 0 1 0-16h16V48a16 16 0 0 1 16-16h40a16 16 0 0 1 16 16v72h16V72a16 16 0 0 1 16-16h40a16 16 0 0 1 16 16v48h16a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}