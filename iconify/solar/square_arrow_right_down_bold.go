package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowRightDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.535 3.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464Zm-5.707 12.114a.75.75 0 0 0 .75-.75v-4.242a.75.75 0 0 0-1.5 0v2.432L9.702 8.64a.75.75 0 0 0-1.06 1.06l4.376 4.377h-2.432a.75.75 0 0 0 0 1.5h4.242Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}