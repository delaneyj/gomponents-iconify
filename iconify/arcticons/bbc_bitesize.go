package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcBitesize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.331h6.858v24.415H5.5zm6.858-5.205h10.833v29.619H12.358zm11.147 1.285l10.363-3.157L42.5 36.587l-10.363 3.158z"/>`),
		g.Group(children),
	)
}