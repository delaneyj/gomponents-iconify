package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronFoldTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 1186l610 610l-145 145l-465-465l-465 465l-145-145l610-610zm0-324L414 252l145-145l465 465l465-465l145 145l-610 610z"/>`),
		g.Group(children),
	)
}