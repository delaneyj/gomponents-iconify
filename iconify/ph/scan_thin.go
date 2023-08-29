package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 40v40a4 4 0 0 1-8 0V44h-36a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4ZM80 212H44v-36a4 4 0 0 0-8 0v40a4 4 0 0 0 4 4h40a4 4 0 0 0 0-8Zm136-40a4 4 0 0 0-4 4v36h-36a4 4 0 0 0 0 8h40a4 4 0 0 0 4-4v-40a4 4 0 0 0-4-4ZM40 84a4 4 0 0 0 4-4V44h36a4 4 0 0 0 0-8H40a4 4 0 0 0-4 4v40a4 4 0 0 0 4 4Zm128 96H88a12 12 0 0 1-12-12V88a12 12 0 0 1 12-12h80a12 12 0 0 1 12 12v80a12 12 0 0 1-12 12Zm4-92a4 4 0 0 0-4-4H88a4 4 0 0 0-4 4v80a4 4 0 0 0 4 4h80a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}