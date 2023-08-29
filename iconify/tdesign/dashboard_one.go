package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v7H2V2Zm2 2v3h5V4H4Zm9-2h9v11h-9V2Zm2 2v7h5V4h-5ZM2 11h9v11H2V11Zm2 2v7h5v-7H4Zm9 2h9v7h-9v-7Zm2 2v3h5v-3h-5Z"/>`),
		g.Group(children),
	)
}