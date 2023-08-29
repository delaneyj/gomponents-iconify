package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2042 1024l-941 941l-90-90l787-787H0V960h1798l-787-787l90-90l941 941z"/>`),
		g.Group(children),
	)
}