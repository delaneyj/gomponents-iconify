package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0Zm9 11C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11Zm-6.914-8L9 11.086l1.414 1.414l-1.5 1.5H14v2H8.914l1.5 1.5L9 18.914L5.086 15ZM10 8h5.086l-1.5-1.5L15 5.086L18.914 9L15 12.914L13.586 11.5l1.5-1.5H10V8Z"/>`),
		g.Group(children),
	)
}