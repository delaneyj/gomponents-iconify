package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneExposureNegOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 18V5h-.3L14 6.7v1.7l3-1.02V18zM4 11h8v2H4z"/>`),
		g.Group(children),
	)
}