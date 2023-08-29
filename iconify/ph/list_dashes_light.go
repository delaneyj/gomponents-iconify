package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDashesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M90 64a6 6 0 0 1 6-6h120a6 6 0 0 1 0 12H96a6 6 0 0 1-6-6Zm126 58H96a6 6 0 0 0 0 12h120a6 6 0 0 0 0-12Zm0 64H96a6 6 0 0 0 0 12h120a6 6 0 0 0 0-12ZM56 58H40a6 6 0 0 0 0 12h16a6 6 0 0 0 0-12Zm0 64H40a6 6 0 0 0 0 12h16a6 6 0 0 0 0-12Zm0 64H40a6 6 0 0 0 0 12h16a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}