package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusPreparingBorderless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><circle cx="8" cy="8" r="2"/><circle cx="14" cy="8" r="2"/><circle cx="2" cy="8" r="2"/></g>`),
		g.Group(children),
	)
}