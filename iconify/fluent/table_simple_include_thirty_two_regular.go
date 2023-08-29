package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleIncludeThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5v7.286a4.742 4.742 0 0 0-2-1.12V7.5A2.5 2.5 0 0 0 24.5 5H17v8.666a4.742 4.742 0 0 0-2 1.12V5H7.5A2.5 2.5 0 0 0 5 7.5V15h9.786a4.742 4.742 0 0 0-1.12 2H5v7.5A2.5 2.5 0 0 0 7.5 27h6.166c.207.76.597 1.443 1.12 2H7.5A4.5 4.5 0 0 1 3 24.5v-17Zm12 10.75A3.25 3.25 0 0 1 18.25 15h7.5A3.25 3.25 0 0 1 29 18.25v7.5A3.25 3.25 0 0 1 25.75 29h-7.5A3.25 3.25 0 0 1 15 25.75v-7.5Z"/>`),
		g.Group(children),
	)
}