package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SessionBorderControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26 8l-1.41 1.41L27.17 12H16a5.967 5.967 0 0 0-4 1.54V6.83l2.59 2.58L16 8l-5-5l-5 5l1.41 1.41L10 6.83V18a5.969 5.969 0 0 0 1.54 4H4.83l2.58-2.59L6 18l-5 5l5 5l1.41-1.41L4.83 24H16a5.99 5.99 0 0 0 4.46-10h6.71l-2.58 2.59L26 18l5-5Zm-6 10a4 4 0 1 1-4-4a4.005 4.005 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}