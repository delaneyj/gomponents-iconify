package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionRecords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCollectionRecords0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 22c0-9.941-8.059-18-18-18S4 12.059 4 22s8.059 18 18 18"/><path stroke-linecap="round" stroke-linejoin="round" d="M33.3 30c-1.822 0-3.3 1.722-3.3 3.846c0 3.845 3.9 7.34 6 8.154c2.1-.813 6-4.31 6-8.154C42 31.722 40.523 30 38.7 30c-1.116 0-2.103.646-2.7 1.634c-.597-.988-1.584-1.634-2.7-1.634Z"/><path fill="#555" d="M22 27a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCollectionRecords0)"/>`),
		g.Group(children),
	)
}