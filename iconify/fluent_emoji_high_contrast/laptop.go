package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="m28.388 21.336l1.397 6.597h-.007l-1.39-6.597Zm-2.542 2.654H6.154l.443-2h18.806l.443 2ZM26 17.75V4.25v13.5ZM7 5h18v12H7V5Z"/>`),
		g.Group(children),
	)
}