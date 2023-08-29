package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStorageCardOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M33.778 4h-18v8h18V4Z"/><path stroke-linecap="round" d="M15.366 4h-4.588a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H34.19"/><path stroke-linecap="round" d="m27.777 20l-8 8.001h10.005l-8.004 8"/></g></mask><path fill="currentColor" d="M0 0h49v48H0z" mask="url(#ipTStorageCardOne0)"/>`),
		g.Group(children),
	)
}