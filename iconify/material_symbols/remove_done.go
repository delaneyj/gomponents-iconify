package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.825 21.8l-6.6-6.6l-2.85 2.85l-5.65-5.65l1.4-1.45l4.25 4.25l1.4-1.4L3.425 3.45L4.825 2l18.4 18.4l-1.4 1.4Zm-15.1-3.75l-5.65-5.65l1.4-1.4l4.25 4.25l1.4 1.4l-1.4 1.4Zm11.3-5.65l-1.4-1.4l4.9-4.9l1.45 1.35l-4.95 4.95Zm-2.85-2.85l-1.4-1.4L15.925 6l1.4 1.4l-2.15 2.15Z"/>`),
		g.Group(children),
	)
}