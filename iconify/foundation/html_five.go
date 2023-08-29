package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m20.706 16.778l5.332 59.802l23.926 6.643l23.993-6.651l5.337-59.793H20.706zM68.172 30.97l-.334 3.718l-.147 1.648H39.624l.67 7.511h26.73l-.179 1.97l-1.724 19.311l-.11 1.238L50 70.526l-.034.011l-15.024-4.171l-1.027-11.517h7.362l.522 5.851l8.168 2.205l.007-.002v-.001l8.181-2.208l.851-9.512h-25.42L31.784 30.97l-.175-1.969h36.739l-.176 1.969z"/>`),
		g.Group(children),
	)
}