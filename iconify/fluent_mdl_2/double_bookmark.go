package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleBookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0v1767l-256-128v385l-640-320l-640 320V256h256V0h1280zm-384 1816V384H384v1432l512-256l512 256zm256-256V128H640v128h896v1240l128 64z"/>`),
		g.Group(children),
	)
}