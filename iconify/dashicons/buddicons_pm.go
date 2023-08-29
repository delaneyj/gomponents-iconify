package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsPm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2c3 0 8 5 8 5v11H2V7s5-5 8-5zm7 14.72l-3.73-2.92L17 11l-.43-.37l-2.26 1.3l.24-4.31l-8.77-.52l-.46 4.54l-1.99-.95L3 11l3.73 2.8l-3.44 2.85l.4.43L10 13l6.53 4.15z"/>`),
		g.Group(children),
	)
}