package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandscapeOrientation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 677v1115H0V128h1499l549 549zm-512-37h293l-293-293v293zm384 1024V768h-512V256H128v1408h1792z"/>`),
		g.Group(children),
	)
}