package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrganizationTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 4.937a2 2 0 1 1 1 0V6h2a1 1 0 0 1 1 1v1.063a2 2 0 1 1-1 0V7h-5v1.063a2 2 0 1 1-1 0V7a1 1 0 0 1 1-1h2V4.937Z"/>`),
		g.Group(children),
	)
}