package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtherAdmissionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.25 18.75q.3 0 .525-.225T16 18q0-.3-.225-.525t-.525-.225q-.3 0-.525.225T14.5 18q0 .3.225.525t.525.225Zm2.75 0q.3 0 .525-.225T18.75 18q0-.3-.225-.525T18 17.25q-.3 0-.525.225T17.25 18q0 .3.225.525t.525.225Zm2.75 0q.3 0 .525-.225T21.5 18q0-.3-.225-.525t-.525-.225q-.3 0-.525.225T20 18q0 .3.225.525t.525.225ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v6.7q-.475-.225-.975-.388T19 11.075V5H5v14h6.05q.075.55.238 1.05t.387.95H5Zm0-3v1V5v6.075V11v7Zm2-1h4.075q.075-.525.238-1.025t.362-.975H7v2Zm0-4h6.1q.8-.75 1.788-1.25T17 11.075V11H7v2Zm0-4h10V7H7v2Zm11 14q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Z"/>`),
		g.Group(children),
	)
}