package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceholderLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 34H48a14 14 0 0 0-14 14v160a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14Zm2 14v153.52L54.49 46H208a2 2 0 0 1 2 2ZM46 208V54.48L201.51 210H48a2 2 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}