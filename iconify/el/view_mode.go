package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1110.645 0L0 1111.963L89.355 1200L1200 89.429L1110.645 0zM0 1.758v281.323h281.25V1.831L0 1.758zm331.421 0v281.323h281.25V1.831l-281.25-.073zM0 338.452v281.25h281.25v-281.25H0zm331.421 0v281.25l281.25-281.25h-281.25zm494.311 297.876L652.295 794.312H1200V636.328H825.732zM614.795 838.33v157.982H1200V838.33H614.795zm0 202.002v157.91H1200v-157.91H614.795z"/>`),
		g.Group(children),
	)
}