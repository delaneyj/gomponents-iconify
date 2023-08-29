package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 64v128a16 16 0 0 1-16 16h-64a24 24 0 0 0-24 24a8 8 0 0 1-16 0a24 24 0 0 0-24-24H32a16 16 0 0 1-16-16V64a16 16 0 0 1 16-16h56a32 32 0 0 1 32 32v88a8 8 0 0 0 16 0V80a32 32 0 0 1 32-32h56a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}