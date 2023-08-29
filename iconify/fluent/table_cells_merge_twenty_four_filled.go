package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsMergeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.25A3.25 3.25 0 0 1 6.25 3h11.5A3.25 3.25 0 0 1 21 6.25v.25H3v-.25ZM3 8v8h18V8H3Zm12.58 3.25l-.89-1.002a.75.75 0 0 1 1.12-.996l1.998 2.247a.748.748 0 0 1 .013.987l-.01.012l-2 2.25a.75.75 0 1 1-1.121-.996l.89-1.002H8.42l.89 1.002a.75.75 0 0 1-1.12.996l-2-2.25a.75.75 0 0 1 0-.996l2-2.25a.75.75 0 1 1 1.12.996l-.89 1.002h7.16ZM3 17.75v-.25h18v.25A3.25 3.25 0 0 1 17.75 21H6.25A3.25 3.25 0 0 1 3 17.75Z"/>`),
		g.Group(children),
	)
}