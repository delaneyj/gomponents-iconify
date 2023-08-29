package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSearchMinus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSearchMinus1" fill="currentColor"><path id="feSearchMinus2" d="m16.325 14.899l5.38 5.38a1.008 1.008 0 0 1-1.427 1.426l-5.38-5.38a8 8 0 1 1 1.426-1.426ZM10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm-3-5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2H7Z"/></g></g>`),
		g.Group(children),
	)
}