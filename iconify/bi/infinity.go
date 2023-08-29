package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.68 5.792L7.345 7.75L5.681 9.708a2.75 2.75 0 1 1 0-3.916ZM8 6.978L6.416 5.113l-.014-.015a3.75 3.75 0 1 0 0 5.304l.014-.015L8 8.522l1.584 1.865l.014.015a3.75 3.75 0 1 0 0-5.304l-.014.015L8 6.978Zm.656.772l1.663-1.958a2.75 2.75 0 1 1 0 3.916L8.656 7.75Z"/>`),
		g.Group(children),
	)
}