package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeOutlineVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 96a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 48a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm0-64a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-48a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm0 144a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 48a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}