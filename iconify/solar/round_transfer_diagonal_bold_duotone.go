package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundTransferDiagonalBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12c0 5.523 4.477 10 10 10a9.954 9.954 0 0 0 5.075-1.382l.007-.004a9.965 9.965 0 0 0 1.238-.864A9.98 9.98 0 0 0 22 12c0-5.523-4.477-10-10-10a9.954 9.954 0 0 0-6.324 2.253A9.98 9.98 0 0 0 2 12Z" opacity=".5"/><path d="M13.129 7.5v2.588L6.917 3.387c-.436.258-.85.548-1.241.867l7.652 8.256a.75.75 0 0 0 1.3-.51V7.5a.75.75 0 1 0-1.5 0Zm-2.709 3.982A.75.75 0 0 0 9.13 12v4.5a.75.75 0 0 0 1.5 0v-2.629l6.446 6.748c.438-.259.854-.55 1.245-.869l-7.9-8.268Z"/></g>`),
		g.Group(children),
	)
}