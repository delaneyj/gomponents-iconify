package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudEditSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a4 4 0 0 0-3.97 3.507A3.25 3.25 0 0 0 4.25 12h1.225c.152-.501.426-.958.798-1.33l4.829-4.83c.248-.247.53-.44.83-.578A4.001 4.001 0 0 0 8 2Zm-1.02 9.377l4.83-4.83a1.87 1.87 0 1 1 2.644 2.646l-4.83 4.829a2.197 2.197 0 0 1-1.02.578l-1.498.374a.89.89 0 0 1-1.079-1.078l.375-1.498a2.18 2.18 0 0 1 .578-1.02Z"/>`),
		g.Group(children),
	)
}