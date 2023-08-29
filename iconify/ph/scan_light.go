package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 40v40a6 6 0 0 1-12 0V46h-34a6 6 0 0 1 0-12h40a6 6 0 0 1 6 6ZM80 210H46v-34a6 6 0 0 0-12 0v40a6 6 0 0 0 6 6h40a6 6 0 0 0 0-12Zm136-40a6 6 0 0 0-6 6v34h-34a6 6 0 0 0 0 12h40a6 6 0 0 0 6-6v-40a6 6 0 0 0-6-6ZM40 86a6 6 0 0 0 6-6V46h34a6 6 0 0 0 0-12H40a6 6 0 0 0-6 6v40a6 6 0 0 0 6 6Zm128 96H88a14 14 0 0 1-14-14V88a14 14 0 0 1 14-14h80a14 14 0 0 1 14 14v80a14 14 0 0 1-14 14Zm2-94a2 2 0 0 0-2-2H88a2 2 0 0 0-2 2v80a2 2 0 0 0 2 2h80a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}