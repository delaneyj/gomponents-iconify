package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><rect width="208" height="160" x="24" y="40" rx="24"/><path d="M160 216H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Z"/></g>`),
		g.Group(children),
	)
}