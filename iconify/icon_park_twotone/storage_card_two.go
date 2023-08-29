package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStorageCardTwo0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="32" height="40" x="8" y="4" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 4v10h20V4"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M14 24h20v20H14z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 32h20m0-3v6m-20-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStorageCardTwo0)"/>`),
		g.Group(children),
	)
}