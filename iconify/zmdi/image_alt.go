package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 320q0 18-12.5 30.5T427 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h384q17 0 29.5 12.5T469 64v256zM160 203l-75 96h299l-96-128l-75 96z"/>`),
		g.Group(children),
	)
}