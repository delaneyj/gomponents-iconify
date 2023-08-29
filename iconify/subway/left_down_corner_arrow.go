package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftDownCornerArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 69.8L442.2 0l-349 349V93.1l-93.1 93L0 512l325.9-.1l93-93.1H163z"/>`),
		g.Group(children),
	)
}