package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M960 1200V0h240v1200H960zM640 300h240v900H640V300zM320 600h240v600H320V600zM0 900h240v300H0V900z"/>`),
		g.Group(children),
	)
}