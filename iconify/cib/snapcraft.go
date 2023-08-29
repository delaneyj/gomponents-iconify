package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.406 17.823V7.589l7.057 3.146zM4.932 31.354l8.656-16.297l3.797 3.792zM0 .646l17.807 6.469v11.307zm29.073 6.469h-10.26l13.188 5.88z"/>`),
		g.Group(children),
	)
}