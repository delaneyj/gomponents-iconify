package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M26.068 10.759H10.965a4.138 4.138 0 0 1 0-8.276h10.759a1.242 1.242 0 0 0 .002-2.482H10.965a6.62 6.62 0 0 0-5.156 10.771l-.01-.013H1.206a1.242 1.242 0 1 0 0 2.484l.037-.001h-.002h15.108a4.138 4.138 0 0 1 0 8.276H5.585A1.242 1.242 0 0 0 5.583 24h10.77A6.615 6.615 0 0 0 21.5 13.229l.01.013h4.597a1.242 1.242 0 1 0-.037-2.483h.002z"/>`),
		g.Group(children),
	)
}