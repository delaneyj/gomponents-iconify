package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#0074BA" d="M3.414 17.414a2 2 0 0 1 0-2.828L14.728 3.272a2 2 0 0 1 2.828 0L28.87 14.586a2 2 0 0 1 0 2.828L17.556 28.728a2 2 0 0 1-2.828 0L3.414 17.414Z"/>`),
		g.Group(children),
	)
}