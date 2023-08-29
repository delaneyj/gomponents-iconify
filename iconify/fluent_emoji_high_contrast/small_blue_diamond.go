package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.636 18.121a3 3 0 0 1 0-4.242l4.243-4.243a3 3 0 0 1 4.242 0l4.243 4.243a3 3 0 0 1 0 4.242l-4.243 4.243a3 3 0 0 1-4.242 0l-4.243-4.243Zm1.414-2.828a1 1 0 0 0 0 1.414l4.243 4.243a1 1 0 0 0 1.414 0l4.243-4.243a1 1 0 0 0 0-1.414l-4.243-4.243a1 1 0 0 0-1.414 0l-4.243 4.243Z"/>`),
		g.Group(children),
	)
}