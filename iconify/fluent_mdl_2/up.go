package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Up(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1875 1037l-787-787v1798H960V250l-787 787l-90-90L1024 6l941 941l-90 90z"/>`),
		g.Group(children),
	)
}