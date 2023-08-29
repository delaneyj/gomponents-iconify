package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeverityLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<circle cx="6" cy="6" r="6" fill="currentColor" fill-rule="evenodd"/>`),
		g.Group(children),
	)
}