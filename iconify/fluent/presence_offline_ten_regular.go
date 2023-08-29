package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceOfflineTenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.854 3.146a.5.5 0 0 1 0 .708L5.707 5l1.147 1.146a.5.5 0 1 1-.708.708L5 5.707L3.854 6.854a.5.5 0 1 1-.708-.708L4.293 5L3.146 3.854a.5.5 0 1 1 .708-.708L5 4.293l1.146-1.147a.5.5 0 0 1 .708 0ZM0 5a5 5 0 1 1 10 0A5 5 0 0 1 0 5Zm5-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}