package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderThumb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1365 341v1366q0 70-27 132t-73 108t-109 74t-132 27q-70 0-132-27t-108-73t-74-109t-27-132V341q0-70 27-132t73-108t109-74t132-27q70 0 132 27t108 73t74 109t27 132z"/>`),
		g.Group(children),
	)
}