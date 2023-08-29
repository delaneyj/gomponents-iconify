package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacetimeVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 134.81v930.381h930.381V673.936L1200 1065.19V134.81L930.381 526.064V134.81H0z"/>`),
		g.Group(children),
	)
}