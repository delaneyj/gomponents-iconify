package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Armrecording(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<circle cx="127" cy="129" r="81" fill="currentColor" fill-rule="evenodd"/>`),
		g.Group(children),
	)
}