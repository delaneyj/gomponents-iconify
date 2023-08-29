package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PavedRoadAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M33 7a1 1 0 0 1 2 0v34a1 1 0 1 1-2 0V7ZM13 7a1 1 0 0 1 2 0v34a1 1 0 1 1-2 0V7Zm10 4a1 1 0 0 1 2 0v4a1 1 0 1 1-2 0v-4Zm0 11a1 1 0 0 1 2 0v4a1 1 0 1 1-2 0v-4Zm0 11a1 1 0 0 1 2 0v4a1 1 0 1 1-2 0v-4Z"/>`),
		g.Group(children),
	)
}