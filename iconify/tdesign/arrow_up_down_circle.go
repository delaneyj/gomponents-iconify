package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-8 6.914L11.086 15l1.414-1.414l1.5 1.5V10h2v5.086l1.5-1.5L18.914 15L15 18.914ZM8 14V8.914l-1.5 1.5L5.086 9L9 5.086L12.914 9L11.5 10.414l-1.5-1.5V14H8Z"/>`),
		g.Group(children),
	)
}