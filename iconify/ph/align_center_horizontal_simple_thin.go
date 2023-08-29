package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontalSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 84h-76V48a4 4 0 0 0-8 0v36H48a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12h76v36a4 4 0 0 0 8 0v-36h76a12 12 0 0 0 12-12V96a12 12 0 0 0-12-12Zm4 76a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V96a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}