package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicnicArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 12a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm2 16h-4.153l-.667-4H24v-2H8v2h2.82l-.667 4H6v2h3.82l-.667 4h2.027l.667-4h8.305l.667 4h2.028l-.667-4H26zm-13.82 0l.666-4h6.307l.666 4z"/>`),
		g.Group(children),
	)
}