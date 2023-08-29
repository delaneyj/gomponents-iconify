package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeBuildingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3v18h-6v-3.5h-2V21H5V3h14m-4 4h2V5h-2v2m-4 0h2V5h-2v2M7 7h2V5H7v2m8 4h2V9h-2v2m-4 0h2V9h-2v2m-4 0h2V9H7v2m8 4h2v-2h-2v2m-4 0h2v-2h-2v2m-4 0h2v-2H7v2m8 4h2v-2h-2v2m-8 0h2v-2H7v2M21 1H3v22h18V1Z"/>`),
		g.Group(children),
	)
}