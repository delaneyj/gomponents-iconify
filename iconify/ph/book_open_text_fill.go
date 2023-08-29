package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenTextFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48h-56a32 32 0 0 0-32 32v88a8 8 0 0 1-16 0V80a32 32 0 0 0-32-32H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h64a24 24 0 0 1 24 24a8 8 0 0 0 16 0a24 24 0 0 1 24-24h64a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16Zm-16 120h-40a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Zm0-32h-40a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Zm0-32h-40a8 8 0 0 1 0-16h40a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}