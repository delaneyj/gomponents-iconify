package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M12.93 32H6.07A2.07 2.07 0 0 1 4 29.93V6.07A2.07 2.07 0 0 1 6.07 4h6.87A2.07 2.07 0 0 1 15 6.07v23.86A2.07 2.07 0 0 1 12.93 32ZM13 6H6v24h7Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M29.93 32h-6.86A2.07 2.07 0 0 1 21 29.93V6.07A2.07 2.07 0 0 1 23.07 4h6.87A2.07 2.07 0 0 1 32 6.07v23.86A2.07 2.07 0 0 1 29.93 32ZM30 6h-7v24h7Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}