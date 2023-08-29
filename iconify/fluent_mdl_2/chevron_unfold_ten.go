package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUnfoldTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M559 815L414 670l610-610l610 610l-145 145l-465-465l-465 465zm930 418l145 145l-610 610l-610-610l145-145l465 465l465-465z"/>`),
		g.Group(children),
	)
}