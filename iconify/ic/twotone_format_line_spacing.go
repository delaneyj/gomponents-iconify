package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneFormatLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 5h12v2H10zm0 12h12v2H10zm-8.5 0L5 20.5L8.5 17H6V7h2.5L5 3.5L1.5 7H4v10zm8.5-6h12v2H10z"/>`),
		g.Group(children),
	)
}