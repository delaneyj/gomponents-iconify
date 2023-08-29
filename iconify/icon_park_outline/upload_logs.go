package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadLogs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M24 44H10a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v18m-4.5 20V31"/><path stroke-linejoin="round" d="m31 34.5l1.5-1.5l3-3l3 3l1.5 1.5"/><path d="M16 16h16m-16 8h8"/></g>`),
		g.Group(children),
	)
}