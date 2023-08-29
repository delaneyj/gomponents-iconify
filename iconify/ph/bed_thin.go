package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 76H20V48a4 4 0 0 0-8 0v160a4 4 0 0 0 8 0v-36h216v36a4 4 0 0 0 8 0v-96a36 36 0 0 0-36-36ZM20 84h80v80H20Zm88 80V84h100a28 28 0 0 1 28 28v52Z"/>`),
		g.Group(children),
	)
}