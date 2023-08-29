package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Purism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 19.588H0V4.412h24zM2.824 16.765h18.352v-9.53H2.824Z"/>`),
		g.Group(children),
	)
}