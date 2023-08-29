package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeOutlineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 98a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18ZM48 98a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm160-48a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}