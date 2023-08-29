package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoneAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.7 18l-5.65-5.65l1.425-1.4l4.25 4.25l1.4 1.4L6.7 18Zm5.65 0L6.7 12.35l1.4-1.425l4.25 4.25l9.2-9.2l1.4 1.425L12.35 18Zm0-5.65l-1.425-1.4L15.875 6L17.3 7.4l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}