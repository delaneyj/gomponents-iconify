package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M572.728 1200V654.546L27.272 1200V0l545.455 545.454V0l600 600l-599.999 600z"/>`),
		g.Group(children),
	)
}