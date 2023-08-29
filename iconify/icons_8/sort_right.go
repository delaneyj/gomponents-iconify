package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 4.594v22.812l1.72-1.687l9-9l.686-.72l-.687-.72l-9-9L12 4.595zm2 4.843L20.563 16L14 22.563V9.438z"/>`),
		g.Group(children),
	)
}