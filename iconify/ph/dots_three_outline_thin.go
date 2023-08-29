package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeOutlineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 100a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm-80-48a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm160-48a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}