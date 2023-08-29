package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeOutlineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 96a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 40a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm80-40a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 40a8 8 0 1 1 8-8a8 8 0 0 1-8 8ZM48 96a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 40a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}