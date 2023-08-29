package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetRoadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm6.177 2.216l-2.393 11.96l-1.96-.391L8.215 5.823l1.96.393Zm5.607-.393l2.393 11.962l-1.962.392l-2.392-11.961l1.961-.393ZM13 7v3h-2V7h2Zm0 4v3h-2v-3h2Zm0 4v3h-2v-3h2Z"/>`),
		g.Group(children),
	)
}