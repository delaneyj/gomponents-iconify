package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m312 61l93 131l-93 131q-13 18-35 18H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h234q22 0 35 18z"/>`),
		g.Group(children),
	)
}