package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatBubbleHistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.959 10.959 0 0 1 1 12Zm12-6.5v6.086L16.414 15L15 16.414l-4-4V5.5h2Z"/>`),
		g.Group(children),
	)
}