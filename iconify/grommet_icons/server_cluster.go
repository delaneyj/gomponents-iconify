package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerCluster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 8h22V1H1v7Zm10-3h1V4h-1v1Zm8 0h1V4h-1v1Zm-4 0h1V4h-1v1Zm-4 7h1v-1h-1v1Zm8 0h1v-1h-1v1Zm-4 0h1v-1h-1v1Zm-4 7h1v-1h-1v1Zm8 0h1v-1h-1v1Zm-4 0h1v-1h-1v1ZM1 15h22V8H1v7Zm0 7h22v-7H1v7Zm20 1H3"/>`),
		g.Group(children),
	)
}